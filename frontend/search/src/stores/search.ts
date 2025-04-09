import { ref } from 'vue';
import { defineStore } from 'pinia';
import type { Track } from '@/stores/models.ts'
import { getSongs } from '@/stores/api.ts'

export const useSearchStore = defineStore('search', () => {
  const results = ref<Track[]>([]);
  const isLoading = ref<boolean>(false);
  const hasSearched = ref<boolean>(false);
  const lastQuery = ref<string>('');

  const search = async (query: string): Promise<void> => {
    isLoading.value = true;
    lastQuery.value = query;
    hasSearched.value = true;

    try {
      const response = await getSongs(query);
      results.value = response.result.results.map(x => x.track);
    } catch (error) {
      console.error('Search error:', error);
      results.value = [];
    } finally {
      isLoading.value = false;
    }
  };

  const clearResults = (): void => {
    results.value = [];
    hasSearched.value = false;
    lastQuery.value = '';
  };

  return {
    results,
    isLoading,
    hasSearched,
    lastQuery,
    search,
    clearResults
  };
});
