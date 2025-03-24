<!-- <template>
    <div id="app">
        <Background />
        <main>
            <SearchBar @search-results="handleResults" />
            <ul>
                <li v-for="(result, index) in searchResults" :key="index">
                    {{ result }}
                </li>
            </ul>
        </main>
    </div>
</template>

<script>
import Background from './components/Background.vue'
import SearchBar from './components/SearchBar.vue'

export default {
    components: { SearchBar, Background },
    data() {
        return {
            searchResults: [] //  JSON
        }
    },
    methods: {
        handleResults(data) {
            this.searchResults = data
        }
    }
}
</script>

<style>
html,
body {
    margin: 0;
    padding: 0;
    height: 100%;
    overflow: hidden;
}

#app {
    position: relative;
    min-height: 100vh;
    overflow: hidden;
}
</style> -->


<template>
    <div id="app">
        <Background />
        <main class="main-container">
            <!-- Окно результатов (слева) -->
            <SearchResults :results="searchResults" />

            <!-- Секция с поисковой строкой (справа) -->
            <div class="search-bar-section">
                <SearchBar @search-results="handleResults" />
            </div>
        </main>
    </div>
</template>

<script>
import Background from './components/Background.vue'
import SearchBar from './components/SearchBar.vue'
import SearchResults from './components/SearchResults.vue'

export default {
    components: {
        Background,
        SearchBar,
        SearchResults
    },
    data() {
        return {
            searchResults: []
        }
    },
    methods: {
        handleResults(data) {
            this.searchResults = data
        }
    }
}
</script>

<style>
html,
body {
    margin: 0;
    padding: 0;
    height: 100%;
    overflow: hidden;
}

#app {
    position: relative;
    min-height: 100vh;
    overflow: hidden;
}

/* Для наглядного разделения слева/справа используем flex */
.main-container {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 2rem;
    /* Расстояние между блоками */
    height: 100vh;
}

/* Можно подстроить стили, чтобы окно результатов и поиск были на одной линии */
.search-bar-section {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
}
</style>