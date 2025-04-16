

/**
 * Helper function to calculate the Euclidean distance between two points.
 * @param {Object} pointA - The first point {lat: number, long: number}.
 * @param {Object} pointB - The second point {lat: number, long: number}.
 * @returns {number} - The Euclidean distance between pointA and pointB.
 */
function calculateDistance(pointA, pointB) {
  const R = 6371; // Earth radius in km
  const lat1 = toRadians(pointA.lat);
  const lat2 = toRadians(pointB.lat);
  const deltaLat = toRadians(pointB.lat - pointA.lat);
  const deltaLong = toRadians(pointB.long - pointA.long);

  const a = Math.sin(deltaLat / 2) * Math.sin(deltaLat / 2) +
            Math.cos(lat1) * Math.cos(lat2) * Math.sin(deltaLong / 2) * Math.sin(deltaLong / 2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

  return R * c; // Return distance in kilometers
}


/**
 * Converts degrees to radians.
 * @param {number} deg - Angle in degrees.
 * @returns {number} - Angle in radians.
 */
const toRadians = (deg) => (deg * Math.PI) / 180;

function drawGraph(ridesArray, origin, destination, ctx, FACTOR) {
  // Initialize graph with nodes (origin, destinations, and waypoints)
  const nodes = [origin, ...ridesArray.map(ride => ride.origin), ...ridesArray.map(ride => ride.destination), destination];
  
  // Create edges (connections between origin, destinations, and waypoints)
  const edges = [];
  ridesArray.forEach(ride => {
      edges.push({ start: ride.origin, end: ride.destination });
  });

  // Draw the nodes (origins, destinations, and waypoints)
  nodes.forEach(node => {
      const x = Math.floor(node.long) * FACTOR;
      const y = Math.floor(node.lat) * FACTOR;
      ctx.beginPath();
      ctx.arc(x, y, 5, 0, Math.PI * 2);
      ctx.fillStyle = "blue";
      ctx.fill();
  });

  // Draw edges (lines between connected nodes)
  edges.forEach(edge => {
      const x1 = Math.floor(edge.start.long) * FACTOR;
      const y1 = Math.floor(edge.start.lat) * FACTOR;
      const x2 = Math.floor(edge.end.long) * FACTOR;
      const y2 = Math.floor(edge.end.lat) * FACTOR;
      
      ctx.beginPath();
      ctx.moveTo(x1, y1);
      ctx.lineTo(x2, y2);
      ctx.strokeStyle = "black";
      ctx.stroke();
  });

  // Draw the path from origin to destination (a simple line to connect start and end)
  const xStart = Math.floor(origin.long) * FACTOR;
  const yStart = Math.floor(origin.lat) * FACTOR;
  const xEnd = Math.floor(destination.long) * FACTOR;
  const yEnd = Math.floor(destination.lat) * FACTOR;

  ctx.beginPath();
  ctx.moveTo(xStart, yStart);
  ctx.lineTo(xEnd, yEnd);
  ctx.strokeStyle = "red"; // Highlight the path in red
  ctx.lineWidth = 2;
  ctx.stroke();

  // Optionally, label the nodes (origins, destinations, waypoints)
  nodes.forEach(node => {
      const x = Math.floor(node.long) * FACTOR;
      const y = Math.floor(node.lat) * FACTOR;
      ctx.font = "12px Arial";
      ctx.fillStyle = "black";
      ctx.fillText(`(${node.lat.toFixed(2)}, ${node.long.toFixed(2)})`, x + 5, y + 5);
  });
}


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

// function drawGraph(ctx, closestRides, pod, FACTOR) {
//   const podX = Math.floor(pod.origin.long) * FACTOR;
//   const podY = Math.floor(pod.origin.lat) * FACTOR;

//   console.log(podX, podY, closestRides);

//   for (let ride of closestRides) {
//       const rideX = Math.floor(ride.origin.long) * FACTOR;
//       const rideY = Math.floor(ride.origin.lat) * FACTOR;

//       console.log(rideX, rideY);

//       // Draw line from pod to ride
//       ctx.beginPath();
//       ctx.moveTo(podX, podY);
//       ctx.lineTo(rideX, rideY);
//       ctx.strokeStyle = "blue";
//       ctx.lineWidth = 2;
//       ctx.stroke();

//       // Label the ride with its ID
//       ctx.fillStyle = "black";
//       ctx.font = "12px Arial";
//       ctx.fillText(ride.rideID, rideX + 5, rideY - 5);
//   }
// }

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
          return { ...ride, distance, bearing: calculateAngleBetweenRides(ride, pod) };
      })
      .sort((a, b) => a.distance - b.distance); 
}

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
  
   const ridesAndPod = await data.json();

   console.log(ridesAndPod)
    
  const ridesArray = ridesAndPod.randomRides;
  const pod = ridesAndPod.pod;
  
  const ranked  = rankRidesByProximityToPod(ridesArray, pod);

  console.log(ridesArray);
  
  console.log(ranked);
  
  const W = ctx.canvas.width,
    H = ctx.canvas.height;

  ctx.setTransform(1, 0, 0, 1, 0, 0);

  ctx.clearRect(0, 0, W, H);

  ctx.setTransform(1, 0, 0, 1, W / 2, H / 2);

  canvas.style.backgroundColor = "white";

  const FACTOR = 3;

  for (let item of ridesArray) {
      // renderPath(item, ctx, FACTOR, "green", "orange")
  }

  // renderPath(pod, ctx, FACTOR)


  drawGraph(ridesArray, pod.origin, pod.destination, ctx,  FACTOR);
}

grapher();
