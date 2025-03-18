<template>
  <div>
    <div id="map"></div>
  </div>
</template>

<script>
import { onMounted, ref, watch } from 'vue';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';

export default {
  props: {
    selectedState: Object, // 传递被选中的州
  },
  setup(props) {
    const map = ref(null);
    const stateLayer = ref(null); // 用于存储当前高亮的州

    onMounted(() => {
      // 创建地图实例
      map.value = L.map('map').setView([37.8, -96], 4); // 默认视角为美国

      // 添加 OpenStreetMap 瓦片层
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
      }).addTo(map.value);
    });

    // 监听 `selectedState` 变化，高亮该州
    watch(() => props.selectedState, (newState) => {
      if (newState && newState.latitude && newState.longitude) {
        // 清除之前的高亮
        if (stateLayer.value) {
          map.value.removeLayer(stateLayer.value);
        }

        // 创建新的高亮层
        stateLayer.value = L.circleMarker([newState.latitude, newState.longitude], {
          radius: 10,
          fillColor: 'red',
          color: 'black',
          weight: 2,
          opacity: 1,
          fillOpacity: 0.6,
        }).addTo(map.value)
          .bindPopup(`<b>${newState.name} (${newState.abbreviation})</b>`)
          .openPopup();

        // 将地图移动到该州
        map.value.setView([newState.latitude, newState.longitude], 6);
      }
    });

    return {};
  },
};
</script>

<style>
#map {
  width: 100%;
  height: 600px; /* 适当放大地图 */
}
</style>
