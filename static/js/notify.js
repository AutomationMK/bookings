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
            <div id="${type}" class="${style} flex items-center gap-2 mx-auto h-fit justify-center ease-in-out duration-500 notify-close overflow-hidden">
            </div>
            `,
        );

        this.noticeEl = this.msgBox.querySelector(`#${this.msgType}`);
    }

    notify(msg, timeout = 0) {
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

        if (timeout > 0) {
            setTimeout(() => {
                this.noticeEl.classList.remove("notify-open");
                this.noticeEl.classList.add("notify-close");
                this.noticeEl.classList.remove(`${this.msgType}-msg`);
                this.noticeEl.innerHTML = "";
            }, timeout);
        }
    }

    addMsgEl(html) {
        this.noticeEl.insertAdjacentHTML("beforeend", `${html}`);
        this.messageEl = this.noticeEl.querySelector("#msg");
    }
}
