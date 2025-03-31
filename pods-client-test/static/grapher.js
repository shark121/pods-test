/**
 * @param {Array<{
 * rideId: string,
 * rideTime: string,
 * rideStatus: string,
 * origin: { lat: number, long: number },
 * destination: { lat: number, long: number },
 * dideCapacity: number,
 * directions: number, // or float, but number is more common in JS
 * dideDistance: number
 * }>} rides - An array of ride objects.
 * @returns {void} - This function does not return a value.
 */

function calculateBearing(origin, destination) {
  const toRadians = (deg) => (deg * Math.PI) / 180;
  const toDegrees = (rad) => (rad * 180) / Math.PI;

  const lat1 = toRadians(origin.lat);
  const lat2 = toRadians(destination.lat);
  const deltaLong = toRadians(destination.long - origin.long);

  const y = Math.sin(deltaLong) * Math.cos(lat2);
  const x = Math.cos(lat1) * Math.sin(lat2) - Math.sin(lat1) * Math.cos(lat2) * Math.cos(deltaLong);

  let bearing = toDegrees(Math.atan2(y, x));
  return (bearing + 360) % 360; // Normalize to 0-360
}

function calculateAngleBetweenRides(ride1, ride2) {
  const bearing1 = calculateBearing(ride1.origin, ride1.destination);
  const bearing2 = calculateBearing(ride2.origin, ride2.destination);

  let angleDifference = Math.abs(bearing1 - bearing2);
  return angleDifference > 180 ? 360 - angleDifference : angleDifference; // Normalize to 0-180
}

function rankRidesByProximityToPod(ridesArray, pod) {
  function getMidpoint(ride) {
      return {
          x: (ride.origin.long + ride.destination.long) / 2,
          y: (ride.origin.lat + ride.destination.lat) / 2
      };
  }

  const podMidpoint = getMidpoint(pod);

  console.log(podMidpoint)

  return ridesArray
      .map(ride => {
          const rideMidpoint = getMidpoint(ride);
          const distance = Math.sqrt(
              (rideMidpoint.x - podMidpoint.x) ** 2 + (rideMidpoint.y - podMidpoint.y) ** 2
          );
          return { rideId: ride.rideId, distance, bearing: calculateAngleBetweenRides(ride, pod) };
      })
      .sort((a, b) => a.distance - b.distance); // Sort in ascending order (closest first)
}


function renderPath(rideObjectType, ctx, FACTOR, originColor, destinationColor ){

    // console.log(rideObjectType.origin.long)
    const originX = Math.floor(rideObjectType.origin.long) * FACTOR;
    const originY = Math.floor(rideObjectType.origin.lat) * FACTOR;
    const destX = Math.floor(rideObjectType.destination.long) * FACTOR;
    const destY = Math.floor(rideObjectType.destination.lat) * FACTOR;
    // console.log(originX, originY, destX, destY);

    ctx.beginPath();
    ctx.arc(originX, originY, 5, 0, Math.PI * 2);
    ctx.fillStyle = originColor ?? "black";
    ctx.fill();

    ctx.beginPath();
    ctx.arc(destX, destY, 5, 0, Math.PI * 2);
    ctx.fillStyle = destinationColor ??  "red";
    ctx.fill();

    ctx.beginPath();
    ctx.moveTo(originX, originY);
    ctx.lineTo(destX, destY);
    ctx.strokeStyle = "black";
    ctx.stroke();

    ctx.font = "14px Arial";
    ctx.fillStyle = "blue";
    const midX = (originX + destX) / 2;
    const midY = (originY + destY) / 2;
    const identifier = rideObjectType.rideId ? rideObjectType.rideId.slice(0,5) : "pod" ;
    ctx.fillText( identifier, midX, midY - 5);

}

export default async function grapher() {
  const canvas = document.getElementById("canvas");

  const ctx = canvas.getContext("2d");

  if (!canvas) {
    console.error("Canvas not found.");
    return;
  }

  if (!ctx) {
    console.error("Failed to get canvas context.");
    return;
  }

  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;

  const data = await fetch("http://localhost:5000", {
    // mode:"no-cors"
  });
  
  console.log(data);
  const ridesAndPod = await data.json();
    
  const ridesArray = ridesAndPod.randomRides;
  const pod = ridesAndPod.pod;
  
  const ranked  = rankRidesByProximityToPod(ridesArray, pod);
  
  console.log(ranked);
  
  const W = ctx.canvas.width,
    H = ctx.canvas.height;

  ctx.setTransform(1, 0, 0, 1, 0, 0);

  ctx.clearRect(0, 0, W, H);

  ctx.setTransform(1, 0, 0, 1, W / 2, H / 2);

  canvas.style.backgroundColor = "white";

  const FACTOR = 3;

  for (let item of ridesArray) {
      renderPath(item, ctx, FACTOR, "green", "orange")
  }

  renderPath(pod, ctx, FACTOR)
}

grapher();
