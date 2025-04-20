<script setup>
import { RouterLink } from 'vue-router';
import Carousel from './Carousel.vue';
</script>

<template>
    <nav>
        <RouterLink to="/home">Home</RouterLink>
    </nav>
    <header class="profile-header flex flex-col-reverse">
        <div class="flex justify-center gap-x-4">
            <!-- a few icons for interaction here will vary depending on state eg add/remove, msg, block -->
            <div class="h-10 w-10 rounded-full bg-gray-300 hover:bg-gray-500"></div>
            <div class="h-10 w-10 rounded-full bg-gray-300 hover:bg-gray-500"></div>
            <div class="h-10 w-10 rounded-full bg-gray-300 hover:bg-gray-500"></div>
        </div>
        <h1 class="self-center text-red-600">(username) profile</h1>
        <!-- maybe add icons for editing username and profile picture -->
        <img
            alt="Vue logo"
            class="self-center"
            src="../assets/logo.svg"
            width="125"
            height="125"
        />
    </header>
    <section>
        <h2 class="text-center">Showcase</h2>
        <Carousel />
    </section>
    <section>
        <h2 class="text-center">Collections</h2>
        <Carousel />
    </section>
</template>

<style scoped></style>
