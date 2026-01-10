export default class Modal {
    modalParent;
    modal;

    constructor() {
        this.modalParent = document.querySelector(".modalbox");

        this.modalParent.insertAdjacentHTML(
            "afterbegin",
            `
            <div class="modal absolute -top-22 hide-modal w-full h-full bg-amber-100/70 backdrop-blur-lg ease-in-out duration-500">
            </div>
            `,
        );

        this.modal = document.querySelector(".modal");
    }

    addHtml(html) {
        this.modal.insertAdjacentHTML("afterbegin", html);
    }

    open() {
        this.modalParent.classList.remove("hide-modalbox");
        this.modalParent.classList.add("show-modalbox");
        this.modal.classList.remove("hide-modal");
        this.modal.classList.add("show-modal");
    }

    close() {
        this.modalParent.classList.remove("show-modalbox");
        this.modalParent.classList.add("hide-modalbox");
        this.modal.classList.remove("show-modal");
        this.modal.classList.add("hide-modal");
    }
}
