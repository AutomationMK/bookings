const nav = document.querySelector(".nav");
const hamburger = document.querySelector(".hamburger");
const slides = document.querySelectorAll(".slide");
const btnLeft = document.querySelector(".slide_btn-left");
const btnRight = document.querySelector(".slide_btn-right");

//////////////////////////////////////
// Mobile Navigation
//////////////////////////////////////
hamburger.addEventListener("click", () => {
  nav.classList.toggle("open");
  hamburger.classList.toggle("close");
});

//////////////////////////////////////
// Slider
//////////////////////////////////////
let curSlide = 0;
const maxSlide = slides.length - 1;

// move to slide with 0 based indexing
const goToSlide = function (index) {
  // check if index is out of bounds
  // if so set the global curSlide to other extreme
  if (index > maxSlide) curSlide = index = 0;
  if (index < 0) curSlide = index = maxSlide;

  slides.forEach(
    (s, i) => (s.style.transform = `translateX(${100 * (i - index)}%)`),
  );
};

// move to next slide
const nextSlide = function () {
  curSlide++;
  goToSlide(curSlide);
};

// move to previous slide
const prevSlide = function () {
  curSlide--;
  goToSlide(curSlide);
};

// initialize the slider
goToSlide(0);

// run button click listeners for slider buttons
btnRight.addEventListener("click", nextSlide);
btnLeft.addEventListener("click", prevSlide);

// run keypad listeners for sliders
document.addEventListener("keydown", function (e) {
  if (e.key === "ArrowRight") {
    nextSlide();
  } else if (e.key === "ArrowLeft") {
    prevSlide();
  }
});
