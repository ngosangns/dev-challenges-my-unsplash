<template>
  <div class="hello">
    <!-- Navbar -->
    <vs-navbar id="top" center-collapsed>
      <template #left>
        <!-- Home button -->
        <vs-button
          gradient
          :active="false"
          @click="resetHome"
        >
          My Unsplash
        </vs-button>

        <!-- Search button -->
        <vs-navbar-item :active="false">
          <vs-input
            id="search-input"
            state="dark"
            type="search"
            placeholder="Search"
            icon-before
            v-on:keyup.enter="search"
            v-model="inSearch"
          >
            <!-- Icon -->
            <template #icon>
              <box-icon name='search'></box-icon>
            </template>
          </vs-input>
        </vs-navbar-item>
      </template>
      <template #right>
        <!-- Add photo button -->
        <vs-button id="add-photo-button" flat :active="true"
          v-on:click="addPhoto()">Add a photo</vs-button>
      </template>
    </vs-navbar>

    <!-- Main -->
    <div id="main">
      <div id="search-results-title" v-if="enableSearchTitle.value">
        <h5>Results for "{{searchText}}"</h5>
      </div>
      <Masonry id="masonry"/>
    </div>

    <AddingPhotoDialog id="dialog" :enableDialog="enableAddingPhotoDialog"/>

    <footer id="footer">
      <span>©2021 ngosangns -
        <a style="color: blue; text-decoration: none;" href="https://devchallenges.io/">devChallenges</a></span>
    </footer>
  </div>
</template>

<script>
import Masonry from "./Masonry/Masonry.vue"
import AddingPhotoDialog from "./AddingPhotoDialog/AddingPhotoDialog.vue"

import axios from "axios"
import host from "@/constants.js"

export default {
  name: 'HelloWorld',
  data: () => ({
    enableAddingPhotoDialog: {
      value: false
    },
    enableDeletingPhotoDialog: {
      value: false
    },
    // Using for state
    enableSearchTitle: {
      value: false,
    },
    inSearch: '', // Using for search input
    searchText: '', // Using for binding
  }),
  components: {Masonry, AddingPhotoDialog},
  props: {
    msg: String
  },
  methods: {
    clearState: function() {
      let states = [
        this.enableAddingPhotoDialog,
        this.enableDeletingPhotoDialog,
        this.enableSearchTitle,
      ]
      for(let state of states) {
        state.value = false
      }
    },
    clearInput: function() {
      this.inSearch = ''
      this.searchText = ''
    },
    resetHome: function() {
      this.clearState()
      this.clearInput()
      this.getData()
    },
    getData: function() {
      // Open loading
      const loading = this.$vs.loading({
        target: document.getElementById("masonry"),
        background: "transparent",
      })
      // Get data
      axios({
        method: 'get',
        url: host+'/get',
      })
      .then(res => {
        this.$store.commit("setImages", res.data.reverse())
      })
      .catch(() => {
        this.$vs.notification({
          color: "danger",
          duration: 3000,
          title: "",
          text: "Error while requesting",
        })
      })
      .finally(() => {
        loading.close()
      })
    },
    addPhoto: function() {
      this.enableAddingPhotoDialog.value = true
    },
    search(e) {
      let searchText = e.target.value

      // Validate
      if(!searchText.length) return

      console.log(searchText)

      // Open loading
      const loading = this.$vs.loading({
        target: document.getElementById("masonry"),
        background: "transparent",
      })
      // Search request
      axios({
        method: 'get',
        url: host+'/search?query='+searchText,
      })
      .then(res => {
        this.$store.commit("setImages", res.data.reverse())
        this.enableSearchTitle.value = true
        this.searchText = searchText
      })
      .catch((err) => {
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
    }
  },
  mounted: function() {
    this.getData()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
#top {
  position: relative;
}

#main {
  position: relative;
  margin: auto;
  margin-top: 1rem;
  width: 96vw;
  min-height: 80vh;
  #search-results-title {

    h5 {
      display: block;
      text-align: left;
      margin: 1rem 0;
    }
  }
}

#footer {
  margin-top: 2rem;
  margin-left: .5rem;
  span {
    display: block;
    text-align: left;
  }
}
</style>
