<template>
  <div>
    <input
      v-model="query"
      @input="searchStates"
      placeholder="Search for a state..."
    />
    <ul v-if="results.length">
      <li
        v-for="state in results"
        :key="state.id"
        @click="selectState(state)"
      >
        {{ state.name }} ({{ state.abbreviation }})
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
import { ref } from 'vue';

export default {
  emits: ['state-selected'], // 让父组件知道用户选择了某个州
  setup(_, { emit }) {
    const query = ref('');
    const results = ref([]);

    const searchStates = async () => {
      if (query.value.length < 1) {
        results.value = [];
        return;
      }

      try {
        const response = await axios.post('http://localhost:8080/graphql', {
          query: `
            query {
              searchStates(query: "${query.value}") {
                id
                name
                abbreviation
                latitude
                longitude
              }
            }
          `,
        });

        results.value = response.data.data.searchStates;
      } catch (error) {
        console.error('Error fetching states:', error);
      }
    };

    const selectState = (state) => {
      emit('state-selected', state);
      query.value = state.name; // 将搜索框更新为选中的州
      results.value = [];
    };

    return { query, results, searchStates, selectState };
  },
};
</script>

<style>
input {
  width: 100%;
  padding: 8px;
  margin-bottom: 5px;
}
ul {
  list-style: none;
  padding: 0;
  margin: 0;
  max-height: 200px;
  overflow-y: auto;
}
li {
  padding: 5px;
  cursor: pointer;
  border-bottom: 1px solid #ccc;
}
li:hover {
  background-color: #f0f0f0;
}
</style>
