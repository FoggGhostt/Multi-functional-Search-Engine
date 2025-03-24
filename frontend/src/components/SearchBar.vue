<template>
    <div class="search-bar-wrapper">
        <input type="text" v-model="query" @keyup.enter="emitSearch" placeholder="search..." />
        <a href="https://github.com/FoggGhostt/Multi-functional-Search-Engine" target="_blank" rel="noopener"
            class="github-link">
            GitHub
        </a>
        <FileUploadButton @file-upload="$emit('file-upload', $event)" />
    </div>
</template>

<script>
import FileUploadButton from './FileUploadButton.vue'

export default {
    components: { FileUploadButton },
    data() {
        return {
            query: ''
        }
    },
    methods: {
        emitSearch() {
            fetch(`http://localhost:8080/api/search?query=${encodeURIComponent(this.query)}`)
                .then(res => res.json())
                .then(data => {
                    this.$emit('search-results', data) // передаём результат родителю
                })
            // .catch(err => console.error("err:", err));
        }
    }
}
</script>

<style scoped>
.search-bar-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100%;
}

input {
    width: 300px;
    height: 30px;
    padding: 0.6rem 1rem;
    border: 2px solid black;
    box-shadow: inset 0 0 0 2px #333;
    border-radius: 999px;
    background-color: transparent;
    color: rgb(8, 8, 8);
    font-size: 1.1rem;
    outline: none;
    transition: box-shadow 0.2s ease;
    font-family: 'Courier New', monospace
}

input::placeholder {
    color: #0e0909;
}

input:focus {
    box-shadow: inset 0 0 0 2px #333, 0 0 0 2px black;
    background-color: rgba(255, 255, 255, 0.05);
}

.github-link {
    position: fixed;
    bottom: 20px;
    right: 20px;
    color: rgb(185, 55, 66);
    text-decoration: none;
    font-weight: bold;
    font-family: 'Courier New', monospace;
    background-color: rgba(8, 8, 8, 0.4);
    padding: 8px 12px;
    border-radius: 999px;
    transition: background-color 0.3s ease;
    z-index: 100;
    /* для уверенности, что кнопка сверху */
}

.github-link:hover {
    background-color: rgba(255, 255, 255, 0.2);
}
</style>