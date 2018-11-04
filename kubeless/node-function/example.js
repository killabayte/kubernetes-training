module.exports = {
  myfunction: function (event, context) {
    console.log(event);
    return "Hello from k8sless function, we are - awesome, as You!!!";
  }
}
