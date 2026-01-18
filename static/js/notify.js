export default class Notice {
    msgBox;
    noticeEl;
    messageEl;
    msgType;

    constructor(msgBoxSel, type, style = "w-screen") {
        this.msgBox = document.querySelector(`${msgBoxSel}`);
        this.msgType = type;

        this.msgBox.insertAdjacentHTML(
            "beforeend",
            `
            <div id="${type}" class="${style} flex mx-auto h-fit justify-center ease-in-out duration-500 notify-close overflow-hidden">
                <h2 class="container px-4 md:text-lg lg:text-xl font-bold text-center"></h2>
            </div>
            `,
        );

        this.noticeEl = this.msgBox.querySelector(`#${type}`);
        this.messageEl = this.noticeEl.querySelector("h2");
    }

    notify(msg) {
        this.noticeEl.classList.remove("notify-close");
        this.noticeEl.classList.add("notify-open");
        if (
            this.msgType === "success" ||
            this.msgType === "warning" ||
            this.msgType === "error"
        ) {
            this.noticeEl.classList.add(`${this.msgType}-msg`);
        }
        this.messageEl.innerHTML = `${msg}`;

        setTimeout(() => {
            this.noticeEl.classList.remove("notify-open");
            this.noticeEl.classList.add("notify-close");
            this.noticeEl.classList.remove(`${this.msgType}-msg`);
            this.messageEl.innerHTML = "";
        }, 5000);
    }
}
