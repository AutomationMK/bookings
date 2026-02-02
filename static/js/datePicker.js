//////////////////////////////////////
// Date Picker
//////////////////////////////////////
export default class Cal {
    curMonth;
    inputBox;
    nextMonth;
    prevMonth;
    realDate;
    startDay;
    endDay;

    // these are global references to bind
    // functions used by the class event handlers
    // I would not have the ability to remove
    // the event listeners otherwise since bind
    // invokes a new reference everytime its called
    handleClickNextCal;
    handleClickPrevCal;
    handleDateSelect;

    constructor(inputBox) {
        this.startDay = null;
        this.endDay = null;

        this.curMonth = new Date();
        this.realDate = new Date();
        this.inputBox = inputBox;

        this.curMonth.setDate(1);

        this.prevMonth = new Date();
        this.#decMonth(this.prevMonth);
        this.nextMonth = new Date();
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

    #incDay(date) {
        date.setDate(date.getDate() + 1);
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
        curCalendar.style.transform = `translateX(0%)`;

        const nextCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.nextMonth.getMonth()}"]`,
        );
        nextCalendar.style.transform = `translateX(${100 * 1}%)`;

        const prevCalendar = this.inputBox.querySelector(
            `.calendar[data-month="${this.prevMonth.getMonth()}"]`,
        );
        prevCalendar.style.transform = `translateX(${100 * -1}%)`;
    }

    selectDate(e) {
        let dateData = e.target.dataset.date;
        let date = null;
        if (!dateData) {
            dateData = e.target.parentElement.dataset.date;
            date = new Date(dateData);
        } else {
            date = new Date(dateData);
        }

        if (this.realDate.getTime() > date.getTime()) {
            // don't allow dates lower than realDate to be selected
            return;
        }

        const inputTextStart =
            this.inputBox.parentElement.querySelector(".start-date-input");

        const inputTextEnd =
            this.inputBox.parentElement.querySelector(".end-date-input");

        let dateDifStart = 0;
        let dateDifEnd = 0;

        if (this.startDay === null) {
            this.startDay = new Date(date);
        } else {
            if (this.startDay.getTime() === date.getTime()) {
                return;
            } else if (
                this.endDay !== null &&
                this.endDay.getTime() === date.getTime()
            ) {
                return;
            } else {
                if (this.endDay === null) {
                    if (this.startDay.getTime() > date.getTime()) {
                        this.endDay = new Date(this.startDay);
                        this.startDay = new Date(date);
                    } else {
                        this.endDay = new Date(date);
                    }
                } else {
                    this.endDay = null;
                    this.startDay = new Date(date);
                    /*
                    dateDifStart = Math.abs(
                        this.startDay.getTime() - date.getTime(),
                    );
                    dateDifEnd = Math.abs(
                        this.endDay.getTime() - date.getTime(),
                    );

                    if (dateDifStart <= dateDifEnd) {
                        this.startDay = date;
                    } else if (dateDifStart > dateDifEnd) {
                        this.endDay = date;
                    }
                    */
                }
            }
        }

        if (this.startDay !== null) {
            inputTextStart.value = this.startDay.toLocaleDateString(
                undefined,
                "YYYY-MM-DD",
            );
        }
        if (this.endDay !== null) {
            inputTextEnd.value = this.endDay.toLocaleDateString(
                undefined,
                "YYYY-MM-DD",
            );
        }

        this.highlightDates();
    }

    highlightDates() {
        this.inputBox.querySelectorAll(".date_box").forEach((el) => {
            el.classList.remove("date_box-active");
            el.classList.add("date_box-inactive");
        });

        let startDayStr = "";
        let endDayStr = "";
        let startDateEl = null;
        let endDateEl = null;

        if (this.startDay === null && this.endDay === null) return;

        if (this.startDay !== null) {
            startDayStr = this.startDay.toLocaleDateString(
                undefined,
                "YYYY-MM-DD",
            );
            startDateEl = this.inputBox.querySelectorAll(
                `.date_box[data-date="${startDayStr}"]`,
            );
        }

        if (this.endDay !== null) {
            endDayStr = this.endDay.toLocaleDateString(undefined, "YYYY-MM-DD");
            endDateEl = this.inputBox.querySelectorAll(
                `.date_box[data-date="${endDayStr}"]`,
            );
        }

        let currentDate = new Date(this.startDay);
        let currentDateStr = "";
        let currentDateEl = null;

        if (endDateEl === null) {
            startDateEl.forEach((el) => {
                el.classList.remove("date_box-inactive");
                el.classList.add("date_box-active");
            });
        } else {
            while (currentDate.getTime() !== this.endDay.getTime()) {
                currentDateStr = currentDate.toLocaleDateString(
                    undefined,
                    "YYYY-MM-DD",
                );
                currentDateEl = this.inputBox.querySelectorAll(
                    `.date_box[data-date="${currentDateStr}"]`,
                );

                currentDateEl.forEach((el) => {
                    el.classList.remove("date_box-inactive");
                    el.classList.add("date_box-active");
                });

                this.#incDay(currentDate);
            }
            endDateEl.forEach((el) => {
                el.classList.remove("date_box-inactive");
                el.classList.add("date_box-active");
            });
        }
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
        this.highlightDates();
    }

    // move to previous cal
    prevCal() {
        if (
            this.prevMonth.getFullYear() < this.realDate.getFullYear() &&
            this.prevMonth.getMonth() > this.realDate.getMonth()
        ) {
            return;
        }

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
        this.highlightDates();
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
            if (
                curDate.getMonth() !== date.getMonth() &&
                curDate.getTime() >= this.realDate.getTime()
            ) {
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
            } else if (
                curDate.getMonth() !== date.getMonth() &&
                curDate.getTime() < this.realDate.getTime()
            ) {
                calendar.insertAdjacentHTML(
                    "beforeend",
                    `
                    <div class="date_box-fade" data-date="${dateStr}">
                        <p class="date_text-fade">
                            ${curDate.getDate()}
                        </p>
                    </div>
`,
                );
            } else if (curDate.getTime() < this.realDate.getTime()) {
                calendar.insertAdjacentHTML(
                    "beforeend",
                    `
                    <div class="date_box-fade" data-date="${dateStr}">
                        <p class="date_text">
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
