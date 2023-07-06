document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('.form');
    const popup = document.getElementById('popup');
    const closeButton = document.getElementById('btn-close')
    const entranceButton = document.querySelector('.entrance-link')
    entranceButton.addEventListener('click', function() {
        console.log("shelk")
        form.style.display = 'block';
    })
    form.addEventListener('submit', function(event) {
        event.preventDefault;
        popup.style.display = 'block';
    })
    closeButton.addEventListener('click', function() {
        popup.style.display = 'none'
    })
})