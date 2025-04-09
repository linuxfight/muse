<template>
  <div class="flex flex-col items-center justify-center min-h-[70vh]">
    <div class="w-full max-w-2xl mx-auto text-center mb-12">
      <h1 class="text-4xl font-bold mb-4 text-yellow-400">Предложить трек</h1>
      <p class="text-gray-400">Найди трек, который ты хочешь предложить на дискотеку</p>
    </div>

    <div class="w-full max-w-2xl mx-auto">
      <div class="relative">
        <input
          v-model="searchQuery"
          @keyup.enter="performSearch"
          type="text"
          placeholder="Toxi$ - Ameli..."
          class="w-full px-5 py-3 pr-12 rounded-lg bg-gray-800 border border-gray-700 text-white focus:outline-none focus:border-transparent"
        />
        <div class="absolute right-3 top-1/2 transform -translate-y-1/2 bg-gray-600 p-2 rounded-lg cursor-pointer" @click="performSearch">
          <button
            class="text-gray-400"
            aria-label="Search"
          >
            <SearchIcon class="h-5 w-5 cursor-pointer" />
          </button>
        </div>
      </div>

      <div v-if="searchStore.isLoading" class="mt-8 text-center">
        <div class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-yellow-400 border-r-transparent"></div>
        <p class="mt-2 text-gray-400">Поиск...</p>
      </div>

      <div v-else-if="searchStore.results != undefined && searchStore.results.length > 0" class="mt-8">
        <h2 class="text-xl font-semibold mb-4 text-yellow-400">Результаты</h2>
        <div class="bg-gray-800 rounded-lg overflow-hidden border border-gray-700">
          <div
            v-for="(result, index) in searchStore.results"
            :key="index"
            class="group p-4 border-b border-gray-700 hover:bg-gray-600 last:border-b-0 flex items-center transition-colors duration-200 cursor-pointer"
            :class="{ 'bg-yellow-400/10 text-yellow-400': selectedSong === result }"
            @click="selectSong(result)"
          >
            <div class="flex-shrink-0 mr-4">
              <img
                :src="getTrackImage(result.coverUri)"
                alt="Thumbnail"
                class="w-16 h-16 object-cover rounded-lg border border-gray-700"
                :class="{ 'border-yellow-400': selectedSong === result }"
              />
            </div>
            <div :class="{ 'text-yellow-400': selectedSong === result }">
              <h3 class="text-lg font-medium text-current">{{ result.title }}</h3>
              <p class="text-gray-400 mt-1" :class="{ 'text-yellow-400': selectedSong === result }">
                {{ result.artists.map(x => x.name).join(' ') }}
              </p>
            </div>
          </div>
        </div>

        <div class="mt-6 flex justify-center">
          <button
            class="px-6 py-3 font-semibold rounded-lg transition-colors duration-200 flex items-center"
            :class="selectedSong ? 'bg-yellow-400 text-gray-900 hover:bg-yellow-500 cursor-pointer' : 'bg-gray-600 text-gray-400 cursor-not-allowed'"
            :disabled="!selectedSong"
            @click="submitSong"
          >
            <span>Предложить трек</span>
          </button>
        </div>
      </div>

      <div v-else-if="searchStore.hasSearched" class="mt-8 text-center">
        <p class="text-gray-400">Ничего не найдено</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useSearchStore } from '@/stores/search.ts';
import { SearchIcon } from 'lucide-vue-next';
import { getTrackImage } from '@/stores/utils.ts';
// import { initData } from '@telegram-apps/sdk-vue';
import type { Track } from '@/stores/models.ts';

const searchQuery = ref('');
const searchStore = useSearchStore();
const selectedSong = ref<Track | null>(null);

const performSearch = async () => {
  if (searchQuery.value) {
    await searchStore.search(searchQuery.value);
    selectedSong.value = null;
  }
};

const selectSong = (song: Track) => {
  if (selectedSong.value === song) {
    selectedSong.value = null;
  } else {
    selectedSong.value = song;
  }
};

const submitSong = () => {
  if (selectedSong.value === null || selectedSong.value === undefined) {
    return;
  }

  alert(selectedSong.value.title);
}
</script>

