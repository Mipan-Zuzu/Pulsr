import axios from "axios";
import { defineStore, storeToRefs } from "pinia";
import { ref } from "vue";
import { Redis } from "./redis";

export const supabaseProject = defineStore("project", () => {
  const redis = Redis()
  const {authKey} = storeToRefs(redis)
  console.log(authKey)
  const projectDetail = ref();
  const projectDetailError = ref<string>("");

  const loadingProject = ref<boolean>(false);
  console.log(authKey.value)

  const allproject = async (value: string
  ) => {
    loadingProject.value = true;
    try {
      const res = await axios.get("http://localhost:3031/v1/supabase/projects", {
    headers: {
        Authorization: value
    },
    withCredentials: true 
});
      console.log(res.data)
      projectDetail.value = res.data
    } catch (err) {
      if (err instanceof Error) {
        projectDetailError.value = err.message;
      }
    } finally {
      loadingProject.value = false;
    }
  };

  return {allproject, projectDetail, loadingProject, projectDetailError};
});

  
  export const supabaseDetail = defineStore("projectdetail", () => {
    const projectData = ref()
    const errProject = ref<String>()
    const sbgsupabase = localStorage.getItem("sbg_supabase")
    const loadingProject = ref<Boolean>(false)
    const allprojectData = async (id: string) => {
      try{
        if (id === "") {
          console.log("id was required")
          return
        }
        loadingProject.value = true
        const res = await axios.get(`http://localhost:3031/v1/supabase/analytics/usage/${id}`,{
        headers: {
            Authorization: `Barer ${sbgsupabase}`
          }
        })
        console.log(res.data)
        projectData.value = res.data
      }catch (err) {
        if (err instanceof Error) {
          errProject.value = err.message
        }
      } finally {
        loadingProject.value = false
      }
    }
    return {allprojectData, projectData, errProject, loadingProject}
  })


