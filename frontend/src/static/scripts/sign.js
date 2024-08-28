/* The code is selecting all elements with the class name "visibility" using
`document.querySelectorAll(".visibility")`. It then iterates over each selected element using
`forEach((eye) => { ... })`. */
const eye = document.querySelectorAll(".visibility");
eye.forEach((eye) => {
  eye.addEventListener("click", async (e) => {
    const input = e.target.previousElementSibling;
    input.type = input.type !== "password" ? "password" : "text";
    e.target.textContent = input.type !== "password" ? "visibility_off" : "visibility";
  });
});
