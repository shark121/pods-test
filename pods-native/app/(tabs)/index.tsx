import React, { useState, useEffect, useRef } from 'react';
import { StyleSheet, View, Text, TouchableOpacity, Dimensions, ActivityIndicator, Alert, TextInput, KeyboardAvoidingView, Platform, ScrollView } from 'react-native';
import MapView, { Marker, Polyline } from 'react-native-maps';
import * as Location from 'expo-location';
import { GooglePlacesAutocomplete } from 'react-native-google-places-autocomplete';
import MapViewDirections from 'react-native-maps-directions';
import { Ionicons } from '@expo/vector-icons';

const { width, height } = Dimensions.get('window');
const GOOGLE_MAPS_API_KEY = process.env.EXPO_PUBLIC_GOOGLE_MAPS_API_KEY || "";
const SERVER_URL = "http://192.168.1.100:5000"; // IMPORTANT: User needs to update this to their local IPv4

type Coordinate = {
  lat: number;
  Lng: number;
};

type PodResponse = {
  message: string;
  rideId: string;
}

export default function HomeScreen() {
  const [origin, setOrigin] = useState<Coordinate | null>(null);
  const [destination, setDestination] = useState<Coordinate | null>(null);
  const [capacity, setCapacity] = useState<number>(2);
  const [selectedTier, setSelectedTier] = useState<string>('Economy');
  const [loading, setLoading] = useState<boolean>(true);
  const [requesting, setRequesting] = useState<boolean>(false);
  const [rideState, setRideState] = useState<'IDLE' | 'WAITING'>('IDLE');
  const [rideId, setRideId] = useState<string | null>(null);
  const [podWaypoints, setPodWaypoints] = useState<Coordinate[]>([]);

  const mapRef = useRef<MapView>(null);

  useEffect(() => {
    (async () => {
      let { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== 'granted') {
        Alert.alert('Permission to access location was denied');
        setLoading(false);
        return;
      }

      let location = await Location.getCurrentPositionAsync({});
      setOrigin({
        lat: location.coords.latitude,
        Lng: location.coords.longitude,
      });
      setLoading(false);
    })();
  }, []);

  const handleRequestRide = async () => {
    if (!origin || !destination) {
      Alert.alert("Missing Locations", "Please ensure both origin and destination are set.");
      return;
    }

    setRequesting(true);
    try {
      // NOTE: Backend expects 'Lng' with capital 'L' 
      const payload = {
        origin: { lat: origin.lat, Lng: origin.Lng },
        destination: { lat: destination.lat, Lng: destination.Lng },
        capacity: capacity
      };

      const response = await fetch(`${SERVER_URL}/api/request-ride`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
      });

      const data: any = await response.json();

      if (response.ok) {
        setRideId(data.rideId);
        setRideState('WAITING');

        if (data.pod && data.pod.podRides) {
          const wps: Coordinate[] = [];
          Object.values(data.pod.podRides).forEach((r: any) => {
            if (r.rideId !== data.rideId) {
              wps.push({ lat: r.origin.lat, Lng: r.origin.Lng });
              wps.push({ lat: r.destination.lat, Lng: r.destination.Lng });
            }
          });
          setPodWaypoints(wps);
        }
      } else {
        Alert.alert("Error", data.message || "Failed to request ride.");
      }
    } catch (error: any) {
      Alert.alert("Network Error", "Could not connect to the pod server. Make sure SERVER_URL matches your local IP.");
      console.error(error);
    } finally {
      setRequesting(false);
    }
  };

  const handleTierSelect = (tier: string, cap: number) => {
    setSelectedTier(tier);
    setCapacity(cap);
  };

  if (loading) {
    return (
      <View style={styles.loadingContainer}>
        <ActivityIndicator size="large" color="#2FCC71" />
        <Text style={styles.loadingText}>Finding your location...</Text>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <MapView
        ref={mapRef}
        style={styles.map}
        initialRegion={{
          latitude: origin?.lat || 37.78825,
          longitude: origin?.Lng || -122.4324,
          latitudeDelta: 0.0922,
          longitudeDelta: 0.0421,
        }}
      >
        {origin && (
          <Marker
            coordinate={{ latitude: origin.lat, longitude: origin.Lng }}
            title="Origin"
            pinColor="black"
          />
        )}

        {destination && (
          <Marker
            coordinate={{ latitude: destination.lat, longitude: destination.Lng }}
            title="Destination"
            pinColor="#2FCC71"
          />
        )}

        {origin && destination && GOOGLE_MAPS_API_KEY !== "" && (
          <MapViewDirections
            origin={{ latitude: origin.lat, longitude: origin.Lng }}
            destination={{ latitude: destination.lat, longitude: destination.Lng }}
            waypoints={podWaypoints.map(wp => ({ latitude: wp.lat, longitude: wp.Lng }))}
            optimizeWaypoints={true}
            apikey={GOOGLE_MAPS_API_KEY}
            strokeWidth={4}
            strokeColor="#2FCC71"
            onReady={(result) => {
              mapRef.current?.fitToCoordinates(result.coordinates, {
                edgePadding: { right: 50, bottom: 350, left: 50, top: 150 },
              });
            }}
          />
        )}
      </MapView>

      {/* Floating Search Bar */}
      {rideState === 'IDLE' && (
        <View style={styles.searchContainer}>
          <GooglePlacesAutocomplete
            placeholder="Where to?"
            fetchDetails={true}
            onPress={(data, details = null) => {
              if (details) {
                setDestination({
                  lat: details.geometry.location.lat,
                  Lng: details.geometry.location.lng
                });
              }
            }}
            query={{
              key: GOOGLE_MAPS_API_KEY,
              language: 'en',
            }}
            styles={{
              container: { flex: 0 },
              textInputContainer: {
                backgroundColor: 'white',
                borderRadius: 25,
                paddingHorizontal: 10,
                alignItems: 'center',
                shadowColor: '#000',
                shadowOffset: { width: 0, height: 2 },
                shadowOpacity: 0.1,
                shadowRadius: 4,
                elevation: 3,
              },
              textInput: {
                height: 48,
                fontSize: 16,
                backgroundColor: 'transparent',
              },
            }}
            renderLeftButton={() => (
              <View style={{ marginLeft: 10 }}>
                <Ionicons name="search" size={20} color="#12141E" />
              </View>
            )}
          />
        </View>
      )}

      {/* Bottom Sheet UI */}
      <View style={styles.bottomSheet}>
        {rideState === 'IDLE' ? (
          <>
            <View style={styles.sheetHeader}>
              <Text style={styles.sheetTitle}>Select Tier</Text>
            </View>

            <View style={styles.tiersContainer}>
              <TouchableOpacity
                style={[styles.tierCard, selectedTier === 'Economy' && styles.tierCardActive]}
                onPress={() => handleTierSelect('Economy', 1)}
              >
                <Ionicons name="car-outline" size={32} color={selectedTier === 'Economy' ? "#2FCC71" : "#12141E"} />
                <Text style={styles.tierName}>Economy</Text>
              </TouchableOpacity>

              <TouchableOpacity
                style={[styles.tierCard, selectedTier === 'Luxury' && styles.tierCardActive]}
                onPress={() => handleTierSelect('Luxury', 2)}
              >
                <Ionicons name="car-sport-outline" size={32} color={selectedTier === 'Luxury' ? "#2FCC71" : "#12141E"} />
                <Text style={styles.tierName}>Luxury</Text>
              </TouchableOpacity>

              <TouchableOpacity
                style={[styles.tierCard, selectedTier === 'Family' && styles.tierCardActive]}
                onPress={() => handleTierSelect('Family', 4)}
              >
                <Ionicons name="bus-outline" size={32} color={selectedTier === 'Family' ? "#2FCC71" : "#12141E"} />
                <Text style={styles.tierName}>Family</Text>
              </TouchableOpacity>
            </View>

            <TouchableOpacity
              style={[styles.btnPrimary, (!origin || !destination) && styles.btnDisabled]}
              disabled={!origin || !destination || requesting}
              onPress={handleRequestRide}
            >
              {requesting ? (
                <ActivityIndicator color="white" />
              ) : (
                <Text style={styles.btnText}>Send Request</Text>
              )}
            </TouchableOpacity>
          </>
        ) : (
          /* Waiting State UI */
          <View style={styles.waitingContainer}>
            <Text style={styles.waitingTitle}>Your Ride</Text>
            <Text style={styles.waitingPrice}>$5.58</Text>
            <Text style={styles.waitingSub}>Searching for pod matches...</Text>
            <Text style={styles.waitingId}>Ride ID: {rideId?.substring(0, 8)}...</Text>

            <ActivityIndicator size="large" color="#2FCC71" style={{ marginVertical: 20 }} />

            <TouchableOpacity
              style={styles.btnPrimary}
              onPress={() => setRideState('IDLE')}
            >
              <Text style={styles.btnText}>Cancel Request</Text>
            </TouchableOpacity>
          </View>
        )}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#fff',
  },
  loadingText: {
    marginTop: 10,
    fontSize: 16,
    color: '#12141E',
    fontWeight: '500'
  },
  map: {
    width: width,
    height: height,
  },
  searchContainer: {
    position: 'absolute',
    top: 60,
    left: 20,
    right: 20,
    zIndex: 1,
  },
  bottomSheet: {
    position: 'absolute',
    bottom: 0,
    left: 0,
    right: 0,
    backgroundColor: 'white',
    borderTopLeftRadius: 30,
    borderTopRightRadius: 30,
    padding: 24,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: -4 },
    shadowOpacity: 0.1,
    shadowRadius: 10,
    elevation: 20,
    minHeight: 280,
  },
  sheetHeader: {
    alignItems: 'center',
    marginBottom: 20,
  },
  sheetTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#12141E',
  },
  tiersContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: 30,
  },
  tierCard: {
    alignItems: 'center',
    justifyContent: 'center',
    padding: 15,
    borderRadius: 15,
    borderWidth: 1,
    borderColor: '#e0e0e0',
    width: '30%',
  },
  tierCardActive: {
    borderColor: '#2FCC71',
    backgroundColor: '#f0fcf5',
  },
  tierName: {
    marginTop: 8,
    fontSize: 14,
    color: '#12141E',
    fontWeight: '600'
  },
  btnPrimary: {
    backgroundColor: '#12141E',
    paddingVertical: 18,
    borderRadius: 30,
    alignItems: 'center',
  },
  btnDisabled: {
    backgroundColor: '#a0a0a0',
  },
  btnText: {
    color: 'white',
    fontSize: 16,
    fontWeight: 'bold',
  },
  waitingContainer: {
    alignItems: 'center',
    paddingVertical: 10,
  },
  waitingTitle: {
    fontSize: 18,
    fontWeight: '600',
    color: '#12141E',
  },
  waitingPrice: {
    fontSize: 36,
    fontWeight: 'bold',
    color: '#12141E',
    marginVertical: 10,
  },
  waitingSub: {
    fontSize: 14,
    color: '#666',
  },
  waitingId: {
    fontSize: 12,
    color: '#999',
    marginTop: 5,
  }
});
