<script setup>
import { onMounted, ref, watch } from 'vue';

const props = defineProps({
  selectedState: Object,
});

const map = ref(null);
const markers = ref([]); // 存储所有标记
const infoWindow = ref(null);

const loadGoogleMaps = () => {
  return new Promise((resolve, reject) => {
    if (window.google && window.google.maps) {
      resolve();
      return;
    }

    const apiKey = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
    if (!apiKey) {
      console.error('Google Maps API Key is missing!');
      reject('API Key is missing');
      return;
    }

    const script = document.createElement("script");
    script.src = `https://maps.googleapis.com/maps/api/js?key=${apiKey}`;
    script.defer = true;
    script.async = true;
    script.onload = resolve;
    script.onerror = () => reject('Google Maps failed to load');
    document.head.appendChild(script);
  });
};

const initMap = () => {
  const mapDiv = document.getElementById("map");
  if (!mapDiv) {
    console.error("Map container #map not found!");
    return;
  }

  map.value = new google.maps.Map(mapDiv, {
    center: { lat: 37.8, lng: -96 },
    zoom: 4,
  });
};

// **刷新页面**
const refreshPage = () => {
  window.location.reload();
};

onMounted(async () => {
  try {
    await loadGoogleMaps();
    initMap();
  } catch (error) {
    console.error("Error loading Google Maps:", error);
  }
});

watch(
  () => props.selectedState,
  (newState) => {
    if (!map.value) return;

    if (newState && newState.latitude && newState.longitude) {
      // **创建新的标记**
      const newMarker = new google.maps.Marker({
        position: { lat: newState.latitude, lng: newState.longitude },
        map: map.value,
        title: newState.name,
      });

      markers.value.push(newMarker); // 存储新标记

      if (!infoWindow.value) {
        infoWindow.value = new google.maps.InfoWindow();
      }

      infoWindow.value.setContent(`<b>${newState.name} (${newState.abbreviation})</b>`);
      infoWindow.value.open(map.value, newMarker);

      // **更新地图中心**
      map.value.setCenter({ lat: newState.latitude, lng: newState.longitude });
      map.value.setZoom(6);
    }
  },
  { immediate: true }
);
</script>

<template>
  <div>
    <button @click="refreshPage" style="margin-bottom: 10px; padding: 8px 12px; background: blue; color: white; border: none; cursor: pointer;">
      Refresh Page
    </button>
    <div id="map" style="width: 100%; height: 500px;"></div>
  </div>
</template>
