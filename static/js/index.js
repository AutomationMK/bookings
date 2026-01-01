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

//////////////////////////////////////
// Date Picker
//////////////////////////////////////
const datePickers = document.querySelectorAll(".date_picker");

class Cal {
    curMonth;
    inputBox;
    nextMonth;
    prevMonth;

    // these are global references to bind
    // functions used by the class event handlers
    // I would not have the ability to remove
    // the event listeners otherwise since bind
    // invokes a new reference everytime its called
    handleClickNextCal;
    handleClickPrevCal;
    handleDateSelect;

    constructor(inputBox) {
        this.curMonth = new Date();
        this.inputBox = inputBox;

        this.curMonth.setDate(1);

        this.prevMonth = new Date();
        this.nextMonth = new Date();
        this.#decMonth(this.prevMonth);
        this.#incMonth(this.nextMonth);

        this.#makeCal(this.prevMonth, "beforeend");
        this.#makeCal(this.curMonth, "beforeend");
        this.#makeCal(this.nextMonth, "beforeend");
        this.updateCal();

        // these are funtion references used by
        // class event listeners
        this.handlePrevCal = this.prevCal.bind(this);
        this.handleNextCal = this.nextCal.bind(this);
        this.handleDateSelect = ((e) => {
            this.selectDate(e);
        }).bind(this);

        this.setDateHandlers();
        this.setButtons();
    }

    #incMonth(date) {
        date.setDate(1);
        let d = date.getDate();
        date.setMonth(date.getMonth() + 1);
        if (date.getDate() !== d) {
            date.setDate(0);
        }
    }

    #decMonth(date) {
        date.setDate(1);
        let d = date.getDate();
        date.setMonth(date.getMonth() - 1);
        if (date.getDate() !== d) {
            date.setDate(0);
        }
    }

    // move to calendar with 0 based indexing
    updateCal() {
        const curCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.curMonth.getMonth()}"]`,
        );

        const prevCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.prevMonth.getMonth()}"]`,
        );

        const nextCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.nextMonth.getMonth()}"]`,
        );

        curCalendar.style.transform = `translateX(0%)`;
        prevCalendar.style.transform = `translateX(${100 * -1}%)`;
        nextCalendar.style.transform = `translateX(${100 * 1}%)`;
    }

    selectDate(e) {
        let date = e.target.dataset.date;
        if (!date) {
            date = e.target.parentElement.dataset.date;
        }

        const inputText =
            this.inputBox.parentElement.querySelector(".date_picker_input");
        inputText.value = date;
    }

    // create date select handlers
    setDateHandlers() {
        const curCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.curMonth.getMonth()}"]`,
        );

        const dateBoxes = curCalendar.querySelectorAll(".date_box");
        dateBoxes.forEach((box) => {
            box.addEventListener("click", this.handleDateSelect);
        });
    }

    // create date select handlers
    resetDateHandlers() {
        const curCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.curMonth.getMonth()}"]`,
        );

        const dateBoxes = curCalendar.querySelectorAll(".date_box");
        dateBoxes.forEach((box) => {
            box.removeEventListener("click", this.handleDateSelect);
        });
    }

    // move to next cal
    nextCal() {
        this.resetDateHandlers();
        this.resetButtons();
        const prevCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.prevMonth.getMonth()}"]`,
        );
        prevCalendar.remove();

        this.#incMonth(this.curMonth);
        this.#incMonth(this.nextMonth);
        this.#incMonth(this.prevMonth);

        this.#makeCal(this.nextMonth, "beforeend");

        this.updateCal();
        this.setDateHandlers();
        this.setButtons();
    }

    // move to previous cal
    prevCal() {
        this.resetDateHandlers();
        this.resetButtons();
        const nextCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.nextMonth.getMonth()}"]`,
        );
        nextCalendar.remove();

        this.#decMonth(this.curMonth);
        this.#decMonth(this.nextMonth);
        this.#decMonth(this.prevMonth);

        this.#makeCal(this.prevMonth, "beforeend");

        this.updateCal();
        this.setDateHandlers();
        this.setButtons();
    }

    #makeCal(date, insertSetting) {
        const monthText = date.toLocaleString("default", { month: "long" });

        this.inputBox.insertAdjacentHTML(
            `${insertSetting}`,
            `
                <div
                    class="calendar" data-month="${date.getMonth()}">
                    <div class="calendar_arrow_box-left">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                            stroke="currentColor" class="calendar_arrow_logo">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5 8.25 12l7.5-7.5" />
                        </svg>
                    </div>
                    <div class="calendar_month_box">
                        <p class="calendar_month_heading">
                            ${monthText} ${date.getFullYear()}
                        </p>
                    </div>
                    <div class="calendar_arrow_box-right">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                            stroke="currentColor" class="calendar_arrow_logo">
                            <path stroke-linecap="round" stroke-linejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
                        </svg>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Sun
                        </p>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Mon
                        </p>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Tue
                        </p>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Wed
                        </p>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Thu
                        </p>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Fri
                        </p>
                    </div>
                    <div class="weekday_box">
                        <p class="date_text">
                            Sat
                        </p>
                    </div>
                </div>`,
        );

        this.fillCalDates(date);
    }

    fillCalDates(date) {
        const monthFirstDate = new Date(
            date.getFullYear(),
            date.getMonth(),
            1,
            0,
            0,
        );

        const monthLastDate = new Date(
            date.getFullYear(),
            date.getMonth() + 1,
            0,
            0,
            0,
        );

        const firstCalDate = new Date(
            date.getFullYear(),
            date.getMonth(),
            -(monthFirstDate.getDay() - 1),
            0,
            0,
        );

        const lastCalDate = new Date(
            date.getFullYear(),
            date.getMonth(),
            monthLastDate.getDate() + 6 - monthLastDate.getDay(),
            0,
            0,
        );

        const numberOfDates =
            monthLastDate.getDate() +
            monthFirstDate.getDay() +
            lastCalDate.getDay() -
            monthLastDate.getDay();

        const calendar = this.inputBox.querySelector(
            `.calendar[data-month="${date.getMonth()}"]`,
        );

        let curDate = new Date(firstCalDate);
        let dateStr = curDate.toLocaleDateString(undefined, "YYYY-MM-DD");
        for (let i = 0; i < numberOfDates; i++) {
            dateStr = curDate.toLocaleDateString(undefined, "YYYY-MM-DD");
            if (curDate.getMonth() !== date.getMonth()) {
                calendar.insertAdjacentHTML(
                    "beforeend",
                    `
                    <div class="date_box" data-date="${dateStr}">
                        <p class="date_text-fade">
                            ${curDate.getDate()}
                        </p>
                    </div>
`,
                );
            } else {
                calendar.insertAdjacentHTML(
                    "beforeend",
                    `
                    <div class="date_box" data-date="${dateStr}">
                        <p class="date_text">
                            ${curDate.getDate()}
                        </p>
                    </div>
`,
                );
            }
            curDate.setDate(curDate.getDate() + 1);
        }
    }

    setButtons() {
        const curCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.curMonth.getMonth()}"]`,
        );
        const rightBtn = curCalendar.querySelector(".calendar_arrow_box-right");
        rightBtn.addEventListener("click", this.handleNextCal);

        const leftBtn = curCalendar.querySelector(".calendar_arrow_box-left");
        leftBtn.addEventListener("click", this.handlePrevCal);
    }

    resetButtons() {
        const curCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.curMonth.getMonth()}"]`,
        );
        const rightBtn = curCalendar.querySelector(".calendar_arrow_box-right");
        rightBtn.removeEventListener("click", this.handleNextCal);

        const leftBtn = curCalendar.querySelector(".calendar_arrow_box-left");
        leftBtn.removeEventListener("click", this.handlePrevCal);
    }
}

datePickers.forEach((datePicker) => {
    const dateInputBox = datePicker.querySelector(".calendar_box");

    let cal = new Cal(dateInputBox);
});
