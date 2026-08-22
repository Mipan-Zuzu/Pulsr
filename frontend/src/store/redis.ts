import axios from "axios";
import { ref} from "vue";
import { defineStore } from "pinia";

export const Redis = defineStore("redis", () => {
    const authKey = ref<string>()
    const errRedis = ref<string>("")

    const redisFind = async() => {
        try {
            const res = await axios.get("http://localhost:3031/v1/redis/dat", {
                withCredentials: true
            })
            authKey.value = res.data.message
            console.log(res.data.message)
        }catch (err) {
            if (err instanceof Error) {
                errRedis.value = err.message
            }
        }
    }
    return {authKey, errRedis, redisFind}
})