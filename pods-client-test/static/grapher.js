/**
 * @param {Array<{
 * RideID: string,
 * RideTime: string,
 * RideStatus: string,
 * Origin: { latitude: number, longitude: number },
 * Destination: { latitude: number, longitude: number },
 * RideCapacity: number,
 * Directions: number, // or float, but number is more common in JS
 * RideDistance: number
 * }>} rides - An array of ride objects.
 * @returns {void} - This function does not return a value.
 */


function renderPath(rideObjectType, ctx, FACTOR, originColor, destinationColor ){

    console.log(rideObjectType.origin.long)
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

  const data = await fetch("http://localhost:8080", {
    // mode:"no-cors"
  });

  const ridesAndPod = await data.json();

  const ridesArray = ridesAndPod.randomRides;
  const pod = ridesAndPod.pod;

  console.log(ridesArray, pod);


  // console.log(ridesArray)

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
