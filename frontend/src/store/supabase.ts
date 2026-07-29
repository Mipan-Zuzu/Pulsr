import axios from "axios";
import { defineStore } from "pinia";
import { ref } from "vue";

export const supabaseProject = defineStore("project", () => {
  const projectDetail = ref();
  const projectDetailError = ref<string>("");

  const loadingProject = ref<boolean>(false);

  const allproject = async (value: string
  ) => {
    loadingProject.value = true;
    try {
      const res = await axios.get("http://localhost:3031/v1/supabase/projects", {
        headers: {
            Authorization: `${value}`  
        }
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
