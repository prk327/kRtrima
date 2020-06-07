//const data = [
//    {
//        width: 200,
//        height: 100,
//        fill: 'purple'
//    },
//    {
//        width: 100,
//        height: 60,
//        fill: 'pink'
//    },
//    {
//        width: 50,
//        height: 30,
//        fill: 'red'
//    }
//];
//
//const svg = d3.select('svg');
//
////Join data to rects
//const rects = svg.selectAll('rect')
//    .data(data);
//
////add attrs to rects already in DOM
//rects.attr('width', d => d.width)
//    .attr('height', d => d.height)
//    .attr('fill', d => d.fill);
//
////append the enter selection to DOM
//rects.enter()
//    .append('rect')
//    .attr('width', d => d.width)
//    .attr('height', d => d.height)
//    .attr('fill', d => d.fill);


//
// start here
//

var gl;
function main() {
  const canvas = document.querySelector("#webgl");
  // Initialize the GL context
  gl = canvas.getContext("webgl");

  // Only continue if WebGL is available and working
  if (gl === null) {
    alert("Unable to initialize WebGL. Your browser or machine may not support it.");
    return;
  }

  // Set clear color to black, fully opaque
  gl.clearColor(0.0, 0.0, 0.0, 1.0);
  // Clear the color buffer with specified clear color
  gl.clear(gl.COLOR_BUFFER_BIT);
    
    let triangle = [
        1.0, -1.0, 0.0,
        0.0, 1.0, 0.0,
        -1.0, -1.0, 0.0
    ];
}

window.onload = main;
