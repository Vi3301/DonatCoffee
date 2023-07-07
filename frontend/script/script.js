document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('.popup_container');
    const popup = document.getElementById('popup');
    const closeButton = document.getElementById('btn-close')
    const entranceButton = document.querySelector('.header_link_entrance')
    entranceButton.addEventListener('click', function() {
        console.log("shelk")
        form.style.display = 'flex';
    })
    form.addEventListener('submit', function(event) {
        event.preventDefault;
        popup.style.display = 'block';
    })
    closeButton.addEventListener('click', function() {
        popup.style.display = 'none'
    })
})