<template>
  <div id="app">
    <h1>US States Map</h1>
    <Typeahead @state-selected="setSelectedState" />
    <MapComponent :selectedState="selectedState" />
  </div>
</template>

<script>
import { ref } from 'vue';
import Typeahead from './components/Typeahead.vue';
import MapComponent from './components/MapComponent.vue';

export default {
  components: {
    Typeahead,
    MapComponent,
  },
  setup() {
    const selectedState = ref(null);

    const setSelectedState = (state) => {
      selectedState.value = state;
    };

    return { selectedState, setSelectedState };
  },
};
</script>

<style>
#app {
  text-align: center;
  padding: 20px;
}
</style>
