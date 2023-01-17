const modal = document.getElementById("myModal");
const span = document.getElementsByClassName("close")[0];

span.onclick = () => {
  modal.style.display = "none";
  location.reload();
};

window.onclick = (event) => {
  if (event.target === modal) {
    modal.style.display = "none";
    location.reload;
  }
};
