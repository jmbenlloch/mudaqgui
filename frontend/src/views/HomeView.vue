<script setup lang="ts">
import MenuBar from '../components/MenuBar.vue';
import ConfigurationView from './ConfigurationView.vue';
import { TabGroup, TabList, Tab, TabPanels, TabPanel } from '@headlessui/vue'
</script>

<template>
  <main>
    <h1 class="text-sm">NEXT Muon DAQ</h1>

    <div class="flex flex-row">
      <div class="w-1/6">
        <MenuBar />
      </div>

      <div class="flex flex-col gap-3">
        <TabGroup>
          <TabList class="tabs">
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Configuration</Tab>
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Plots</Tab>
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Tab 3</Tab>
          </TabList>
          <TabPanels>
            <TabPanel>
              <ConfigurationView />
            </TabPanel>
            <TabPanel>Content 2</TabPanel>
            <TabPanel>Content 3</TabPanel>
          </TabPanels>
        </TabGroup>
      </div>
    </div>
  </main>
</template>
