<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router';
import Carousel from './Carousel.vue';
import { onMounted } from 'vue';

const trendingCards = ref([]);

onMounted(() => {
    fetch('http://localhost:8080/cards')
        .then((res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                throw new Error('Failed to fetch trending cards');
            }
        })
        .then((data) => {
            trendingCards.value = data;
        })
        .catch((e) => console.error(e.message));
});
const searchString = ref('');
const searching = ref(false);
const searchResults = ref([]);

function search() {
    searching.value = true;
    fetch('http://localhost:8080/search', {
        method: 'POST',
        body: JSON.stringify({ name: searchString.value }),
    })
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else if (res.status == 404) {
                throw new Error('The requested resource was not found on the server.');
            } else {
                throw new Error('Unexpected response ' + res.statusText);
            }
        })
        .then((data) => {
            console.log(data);
            searchResults.value = data;
        })
        .catch((e) => console.error(e.message))
        .finally(() => {
            searchString.value = '';
            searching.value = false;
        });
}
</script>

<template>
    <div>
        <div class="flex items-center justify-between gap-x-3">
            <header class="flex items-center justify-center">
                <img
                    alt="Vue logo"
                    class="logo"
                    src="../assets/logo.svg"
                    width="125"
                    height="125"
                />
                <h1 class="h-min text-nowrap text-red-600">Pok&#233;mon Collectors</h1>
            </header>
            <nav>
                <h1 class="hover:text-gray-600"><RouterLink to="/">X</RouterLink></h1>
                <!-- hamburger menu -->
            </nav>
        </div>
        <section>
            <RouterLink to="/profile">Profile</RouterLink>
        </section>
        <section>
            <input
                type="text"
                class="bg-dark text-dark rounded-md pl-1"
                placeholder="Search for cards"
                v-model="searchString"
            />
            <button
                class="btn-white text-dark ml-2"
                @click="search"
            >
                {{ searching ? 'Searching...' : 'Search' }}
            </button>
        </section>
        <section v-if="searchResults.length">
            <Carousel :preview="searchResults" />
        </section>
        <section v-if="trendingCards.length">
            <h2 class="text-center">Trending Cards</h2>
            <Carousel :preview="trendingCards" />
        </section>
        <section>
            <!-- <h2 class="text-center">Trending Collections</h2> -->
            <!-- <Carousel /> -->
        </section>
    </div>
</template>

<style scoped></style>
