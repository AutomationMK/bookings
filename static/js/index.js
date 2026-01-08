const nav = document.querySelector(".nav");
const hamburger = document.querySelector(".hamburger");
const slider = document.querySelector(".slider");
const slides = document.querySelectorAll(".slide");
const btnLeft = document.querySelector(".slide_btn-left");
const btnRight = document.querySelector(".slide_btn-right");
const dotContainer = document.querySelector(".dots");

//////////////////////////////////////
// Mobile Navigation
//////////////////////////////////////
if (hamburger !== null) {
    hamburger.addEventListener("click", () => {
        nav.classList.toggle("open");
        hamburger.classList.toggle("close");
    });
}

//////////////////////////////////////
// Slider
//////////////////////////////////////
let curSlide = 0;
const maxSlide = slides.length - 1;

// create dot slider elements based on amount of slides
const createDots = function () {
    slides.forEach(function (_, i) {
        dotContainer.insertAdjacentHTML(
            "beforeend",
            `<button class="dots_dot" data-slide="${i}"></button>`,
        );
    });
};

// activate a dot to show which slide is active
const activateDot = function (index) {
    document
        .querySelectorAll(".dots_dot")
        .forEach((dot) => dot.classList.remove("dots_dot-active"));

    document
        .querySelector(`.dots_dot[data-slide="${index}"]`)
        .classList.add("dots_dot-active");
};

// move to slide with 0 based indexing
const goToSlide = function (index) {
    // check if index is out of bounds
    // if so set the global curSlide to other extreme
    if (index > maxSlide) curSlide = index = 0;
    if (index < 0) curSlide = index = maxSlide;

    activateDot(index);
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

// implement touch swiping of slider
var xDown = null;

const handleTouchStart = function (e) {
    const firstTouch = e.touches[0];
    xDown = firstTouch.clientX;
};

const handleTouchSlide = function (e) {
    if (!xDown) {
        return;
    }

    var xUp = e.touches[0].clientX;

    var xDiff = xDown - xUp;

    if (xDiff > 5 && curSlide < maxSlide) {
        // right slide
        nextSlide();
    } else if (xDiff < -5 && curSlide > 0) {
        // left slide
        prevSlide();
    }
    // reset global value
    xDown = null;
};

// validate if a slider is on the page
if (slider !== null) {
    // initialize the slider and dots
    createDots();
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

    // run listener for click on dots
    dotContainer.addEventListener("click", function (e) {
        if (e.target.classList.contains("dots_dot")) {
            curSlide = Number(e.target.dataset.slide);
        }
        goToSlide(curSlide);
    });

    slider.addEventListener("touchstart", handleTouchStart);
    slider.addEventListener("touchmove", handleTouchSlide);
}
