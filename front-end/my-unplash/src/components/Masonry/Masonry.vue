<template>
    <div>
        <!-- Masonry -->
        <vue-masonry-wall id="masonry" :options="{width: 300, padding: 12}"
            :items="masonryItems">
            <template v-slot:default="{item}">
                <div class="item" v-if="item">
                    <div class="button-on-hover-container">
                        <vs-button
                            danger
                            border
                            :active="false"
                            class="button-on-hover"
                            v-on:click="onDeletingPhoto(item)"
                        >
                            delete
                        </vs-button>
                    </div>
                    <div class="photo-container">
                        <img :src="item.data.link" style="width: 100%">
                    </div>
                    <div class="label-on-hover-container">
                        <span class="label-on-hover">
                            #{{item.data.label}}
                        </span>
                    </div>
                </div>
            </template>
        </vue-masonry-wall>
        <div>
            <vs-pagination v-model="currentPage.value" :length="pageLength" not-arrows
                @input="handlePageChange" v-if="this.$store.state.images.length"/>
        </div>
        <DeletingPhotoDialog :enableDialog="enableDeletingPhotoDialog" :photo="deletedPhoto"/>
    </div>
</template>

<script>
import VueMasonryWall from "vue-masonry-wall"
import DeletingPhotoDialog from "@/components/Masonry/DeletingPhotoDialog/DeletingPhotoDialog.vue"

export default {
    name: "Masonry",
    components: {VueMasonryWall, DeletingPhotoDialog},
    data: () => ({
        enableDeletingPhotoDialog: {
            value: false,
        },
        deletedPhoto: {
            value: null
        },
        // Pagination
        currentPage: {
            value: 1,
        },
        itemsPerPage: 20,
    }),
    watch: {
        masonryItems: function() {
            // Go to newer page if current page is null
            if(this.currentPage.value > this.pageLength && this.pageLength > 0) {
                this.currentPage.value = this.pageLength
            }
        }
    },
    computed: {
        pageLength: function() {
            return Math.ceil(this.$store.state.images.length/this.itemsPerPage)
        },
        masonryItems: function() {
            let startIndex = (this.currentPage.value-1)*this.itemsPerPage
            let endIndex = startIndex + this.itemsPerPage
            return this.$store.state.images.slice(startIndex, endIndex)
        },
    },
    methods: {
        handlePageChange: function(val) {
            this.currentPage.value = val
        },
        onDeletingPhoto: function(photo) {
            this.deletedPhoto.value = photo
            this.enableDeletingPhotoDialog.value = true
        }
    },
}
</script>

<style lang="scss" scoped>
.item {
    display: grid;
    grid-template-rows: auto max-content auto;
    background-color: black;
    border-radius: 1rem;
    overflow: hidden;
    .button-on-hover-container {
        display: none;
        justify-content: flex-end;
        .button-on-hover {
            margin: .6rem;
            position: absolute;
        }
    }
    .label-on-hover-container {
        display: none;
        justify-content: flex-start;
        align-items: flex-end;
        .label-on-hover {
            color: white;
            padding: 1rem;
            font-size: 1.2rem;
            position: absolute;
        }
    }
    .photo-container,.button-on-hover-container, .label-on-hover-container {
        transition: all .2s ease-in-out;
    }
}
.item:hover {
    .photo-container {
        opacity: .5;
    }
    .button-on-hover-container, .label-on-hover-container {
        display: flex;
    }
    .photo-container,.button-on-hover-container, .label-on-hover-container {
        transition: all .2s ease-in-out;
    }
}
</style>