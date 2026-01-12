export default class Notice {
    msgBox;
    noticeEl;
    messageEl;

    constructor(msgBoxSel, id) {
        this.msgBox = document.querySelector(`${msgBoxSel}`);

        this.msgBox.insertAdjacentHTML(
            "beforeend",
            `
            <div id="${id}" class="flex w-screen h-fit justify-center ease-in-out duration-500 notify-close overflow-hidden">
                <h2 class="text-2xl font-bold"></h2>
            </div>
            `,
        );

        this.noticeEl = this.msgBox.querySelector(`#${id}`);
        this.messageEl = this.noticeEl.querySelector("h2");
    }

    notify(msg, type) {
        this.noticeEl.classList.remove("notify-close");
        this.noticeEl.classList.add("notify-open");
        if (type === "success" || type === "warning" || type === "error") {
            this.noticeEl.classList.add(`${type}-msg`);
        }
        this.messageEl.innerHTML = `${msg}`;

        setTimeout(() => {
            this.noticeEl.classList.remove("notify-open");
            this.noticeEl.classList.add("notify-close");
            this.noticeEl.classList.remove(`${type}-msg`);
            this.messageEl.innerHTML = "";
        }, 5000);
    }
}
