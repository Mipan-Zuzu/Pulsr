import { defineStore } from "pinia"
import {ref} from "vue"

export const useNavSidebar = defineStore('shared', () => {
    const myState = ref("")

    const setNav = (value: string) => {
        myState.value = value
    }

    return {myState, setNav}
})