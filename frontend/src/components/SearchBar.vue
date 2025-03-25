<template>
    <div class="search-bar-wrapper">
        <!-- Само "поисковое окно" с input и кнопкой внутри -->
        <div class="search-bar">
            <input type="text" v-model="query" @keyup.enter="emitSearch" placeholder="search..." />
            <!-- Кнопка с лупой внутри того же контейнера -->
            <button @click="emitSearch" title="Поиск">🔍</button>
        </div>

        <!-- Ссылка на GitHub -->
        <a href="https://github.com/FoggGhostt/Multi-functional-Search-Engine" target="_blank" rel="noopener"
            class="github-link">
            GitHub
        </a>

        <!-- Кнопка загрузки файлов -->
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
                .catch(err => console.error('Ошибка поиска:', err))
        }
    }
}
</script>

<style scoped>
/* Контейнер, в котором располагается "поисковое окно", GitHub-ссылка и кнопка загрузки */
.search-bar-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100%;
}

/* Сам блок, объединяющий input и кнопку лупы */
.search-bar {
    display: flex;
    align-items: center;
    border: 2px solid black;
    box-shadow: inset 0 0 0 2px #333;
    border-radius: 999px;
    background-color: transparent;
    transition: box-shadow 0.2s ease, transform 0.2s ease;
    padding: 0.3rem 0.6rem;
    /* Немного отступов, чтобы текст и кнопка не прилипали к краям */
}

/* Эффект при наведении на всё "поисковое окно" */
.search-bar:hover {
    box-shadow: inset 0 0 0 2px #333, 0 0 0 2px black;
    transform: scale(1.02);
    /* Лёгкое увеличение */
}

input {
    width: 300px;
    height: 35px;
}

/* Поле ввода */
.search-bar input {
    flex: 1;
    /* Заставляет input занимать всё доступное пространство, а кнопка будет прижата справа */
    border: none;
    outline: none;
    background: transparent;
    color: rgb(8, 8, 8);
    font-size: 1.1rem;
    font-family: 'Courier New', monospace;
    margin-right: 8px;
    /* Отступ между input и кнопкой */
}

.search-bar input::placeholder {
    color: #0e0909;
    font-size: 14px;
}

/* Кнопка с лупой */
.search-bar button {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: rgba(7, 7, 7, 0.4);
    border: 2px solid black;
    color: #fff; /* для хорошей читаемости на темном фоне */
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    transition: background-color 0.3s ease;
}

.search-bar button:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

/* Ссылка на GitHub */
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
    /* Чуть выше, чтобы ссылка не пряталась за другими элементами */
}

.github-link:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

</style>


