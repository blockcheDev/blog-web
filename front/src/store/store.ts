import { reactive, ref } from "vue";

export const drawerMenu = reactive({
    isOpen: false,
    open() {
        this.isOpen = true
    },
    close() {
        this.isOpen = false
    }
})

export const loginDialog = reactive({
    isOpen: false,
    open() {
        this.isOpen = true
    },
    close() {
        this.isOpen = false
    }
})