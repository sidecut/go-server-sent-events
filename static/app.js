// @ts-check

document.addEventListener("DOMContentLoaded", function () {
  const eventSource = new EventSource("./progress");
  eventSource.addEventListener("onProgress", (event) => {
    const data = JSON.parse(event.data);
    const { progressPercentage /*, message*/ } = data;
    console.log(progressPercentage /*, message*/);

    const progressBar = document.getElementById("progressBar");
    progressBar && (progressBar.style.width = `${5 * progressPercentage}%`);
  });

  eventSource.addEventListener("onComplete", (event) => {
    eventSource.close();
    console.log("onComplete");
  });
});
