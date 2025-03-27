<template>
    <div id="app">
        <Background />
        <main class="main-container">
            <SearchResults :results="searchResults" />
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

.main-container {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 2rem;
    height: 100vh;
}

.search-bar-section {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
}
</style>