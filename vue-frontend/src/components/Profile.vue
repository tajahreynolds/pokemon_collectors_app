<script setup>
import { RouterLink } from 'vue-router';

</script>

<template>
    <nav>
        <RouterLink to="/home">Home</RouterLink>
    </nav>
    <header class="profile-header flex flex-col-reverse">
        <div class="flex justify-center gap-x-4">
            <!-- a few icons for interaction here will vary depending on state eg add/remove, msg, block -->
            <div class="w-10 h-10 bg-gray-300 rounded-full hover:bg-gray-500"></div>
            <div class="w-10 h-10 bg-gray-300 rounded-full hover:bg-gray-500"></div>
            <div class="w-10 h-10 bg-gray-300 rounded-full hover:bg-gray-500"></div>
        </div>
        <h1 class="text-red-600 self-center">(username) profile</h1>
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
        <!-- image carousel with controls -->
        <div class="flex overflow-x-auto gap-x-3">
            <div v-for="n in 3" key="n" class="min-w-1/2 text-center h-80 bg-dark rounded-lg">
                <h3 class="text-dark">Placeholder</h3>
            </div>
        </div>
    </section>
    <section>
        <h2 class="text-center">Collections</h2>
        <!-- image carousel with controls -->
        <div class="flex overflow-x-auto gap-x-3">
            <div v-for="n in 3" key="n" class="min-w-1/2 text-center h-80 bg-dark rounded-lg">
                <h3 class="text-dark">Placeholder</h3>
            </div>
        </div>
    </section>
</template>

<style scoped>
</style>
