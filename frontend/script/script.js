document.addEventListener("DOMContentLoaded", function () {
  const popupContainer = document.querySelector(".popup_container");
  const submitButtonForm = document.querySelector(".form_btn-submit")
  const formContainer = document.querySelector(".form_container")
  const form = document.querySelector(".form")
  const popup = document.getElementById("popup");
  const closeButton = document.getElementById("btn-close");
  const entranceButton = document.querySelector(".header_link_entrance");
  const closeIconForm = document.querySelector(".popup_close-form-icon")
  function openForm() {
    popupContainer.style.display = "flex";
  }

  function closeForm() {
    popupContainer.style.display = "none";
    
  }
  function submitForm(event) {
    event.preventDefault();
    popup.style.display = "block"
  }
  
  function submitUser() {
    fetch('http://localhost:8080/api/greeting')
    .then(response => response.text())
    .then(data => console.log(data))
    .catch((error) => {
        console.error('Error', error)
    })
  }
   closeButton.addEventListener('click', function() {
   popup.style.display = 'none'
  })
  submitButtonForm.addEventListener("click", submitUser)
  closeIconForm.addEventListener("click", closeForm)
  form.addEventListener('submit', submitForm)
  entranceButton.addEventListener("click", openForm);
  closeButton.addEventListener("click", closeForm)

});


