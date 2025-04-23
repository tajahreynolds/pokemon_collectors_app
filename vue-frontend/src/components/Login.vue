<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

function gotoHomePage() {
    router.push({ name: 'Home' });
}

const loginEmail = ref('');

function login(e) {
    e.preventDefault();
    e.stopPropagation();

    fetch('http://localhost:8080/login', {
        method: 'POST',
        body: JSON.stringify({ email: loginEmail.value }),
    })
        .then((res) => res.json())
        .then((data) => {
            if (data.message == 'success') {
                gotoHomePage();
            } else {
                throw new Error(
                    'Fail - Account not found ? Should be treated the same as account found (next step: click confirmation link)'
                );
            }
        })
        .catch((err) => {
            console.error(err);
            alert('There was an error trying to login to your account. Please try again!');
        });
}
</script>

<template>
    <div>
        <header class="flex flex-row items-center">
            <img
                alt="Vue logo"
                class="logo"
                src="../assets/logo.svg"
                width="125"
                height="125"
                @click="gotoHomePage()"
            />
            <h1 class="text-red-600">Pok&#233;mon Collectors</h1>
        </header>

        <section class="flex flex-col items-center gap-4 p-4">
            <h1 class="text-center min-[551px]:max-[976px]:max-w-[532px]">
                Log in to your account to gain access to all features!
            </h1>
            <form
                class="w-3/4 min-w-[277px] rounded-lg bg-white p-4 min-[610px]:max-[786px]:min-w-[523px]"
                @submit="login"
            >
                <div class="pb-4">
                    <label
                        for="email"
                        class="text-dark block"
                        >Email Address:</label
                    >
                    <input
                        type="email"
                        id="email"
                        name="email"
                        class="w-full rounded border px-3 py-2 text-gray-700"
                        placeholder="Enter your email"
                        v-model="loginEmail"
                        required
                    />
                </div>
                <div class="text-dark flex items-center justify-between">
                    <span class="pr-9"
                        >Need to create an account?
                        <RouterLink to="Register">Register</RouterLink></span
                    >
                    <button
                        class="btn-red mr-4"
                        type="submit"
                    >
                        Login
                    </button>
                </div>
            </form>
        </section>
    </div>
</template>

<style scoped></style>
