<template>
    <vs-dialog :not-close="true" v-model="enableDialog.value" overflow-hidden prevent-close>
      <template #header>
        <h4>Add a new photo</h4>
      </template>

      <div id="dlg-form">
        <vs-input class="dlg-input"
          state="dark"
          label="Label"
          v-model="photo.label"
          placeholder="nature"
        >
          <!-- Message -->
          <template #message-danger>
            {{message.label}}
          </template>
        </vs-input>
        <vs-input class="dlg-input"
          state="dark"
          label="Photo URL"
          v-model="photo.link"
          placeholder="https://images.unsplash.com/photo-16117..."
        >
          <!-- Message -->
          <template #message-danger>
            {{message.link}}
          </template>
        </vs-input>
      </div>

      <template #footer>
        <div id="dlg-footer">
          <vs-button dark flat :active="false" v-on:click="hideDialog">
            Cancel
          </vs-button>
          <vs-button flat :active="true" v-on:click="onSubmit">
            Submit
          </vs-button>
        </div>
      </template>
    </vs-dialog>
</template>
<script>
import axios from "axios"
import host from "@/constants.js"

export default {
    name: "AddingPhotoDialog",
    props: ["enableDialog"],
    data: () => ({
        photo: {
          label: "",
          link: "",
        },
        message: {
          label: "",
          link: "",
        },
    }),
    methods: {
      hideDialog: function() {
        this.clearForm()
        this.enableDialog.value = false
      },
      clearForm: function() {
        for(let key in this.photo) {
          this.photo[key] = ""
        }
        for(let key in this.message) {
          this.message[key] = ""
        }
      },
      onSubmit: function() {
        // Validate
        if(!this.photo.label.match(/^[0-9a-zA-Z_]+$/)) {
          this.message.label = "The field only contains numbers, digits, and underscore"
          return
        } else this.message.label = ""
        if(!this.photo.link.match(/^https:\/\/images.unsplash.com\//)) {
          this.message.link = "The field must be Unsplash's image link"
          return
        } else this.message.link = ""

        // Create form data
        let formData = new FormData()
        for(let key in this.photo) {
          formData.append(key, this.photo[key])
        }
        let loading = this.createLoading()

        // Send
        axios({
          method: 'post',
          url: host+'/create',
          data: formData,
        }).then(res => {
          // Notification
          this.$vs.notification({
            color: "success",
            duration: 3000,
            title: "",
            text: "Upload success!"
          })
          this.$store.commit("unshiftImage", res.data)
          this.clearForm()
          this.enableDialog.value = false
        })
        .catch(err => {
          // Check response message
          if (err.response) {
            if([501].includes(err.response.status)) {
              this.$vs.notification({
                color: "danger",
                duration: 3000,
                title: '',
                text: err.response.data.message
              })
            }
          }
          else {
            this.$vs.notification({
              color: "danger",
              duration: 3000,
              title: "",
              text: "Error while requesting",
            })
          }

          throw new Error(err)
        })
        .finally(() => {
          loading.close()
        })
      },
      createLoading: function() {
        const loading = this.$vs.loading()
        // loading.close()
        return loading
      },
    },
}
</script>
<style lang="scss">
#dlg-form {
    width: 100%;
    .dlg-input {
        width: 100%;
        margin-top: 1.75rem;
        .vs-input {
            width: 100%;
        }
    }
}

#dlg-footer {
    display: flex;
    justify-content: flex-end;
}
</style>