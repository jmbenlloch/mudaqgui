<script setup lang="ts">
import MenuBar from '../components/MenuBar.vue';
import ConfigurationView from './ConfigurationView.vue';
import EventsView from './EventsView.vue';
import ChannelView from './ChannelView.vue';
import AllChannelsView from './AllChannelsView.vue';
import Rate from '../components/Rate.vue';
import { TabGroup, TabList, Tab, TabPanels, TabPanel } from '@headlessui/vue'
</script>

<template>
  <main>
    <div class="flex flex-row">
      <div class="w-2/12 bg-base-200">
        <MenuBar />
      </div>

      <div class="flex flex-col gap-2">
        <TabGroup>
          <TabList class="tabs">
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Configuration</Tab>
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Rates</Tab>
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Plots</Tab>
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">Single channel</Tab>
            <Tab class="tab tab-lg tab-lifted ui-selected:tab-active">All channels</Tab>
          </TabList>
          <TabPanels>
            <TabPanel>
              <ConfigurationView />
            </TabPanel>
            <TabPanel>
              <Rate />
            </TabPanel>
            <TabPanel>
              <EventsView />
            </TabPanel>
            <TabPanel>
              <ChannelView />
            </TabPanel>
            <TabPanel>
              <AllChannelsView />
            </TabPanel>
          </TabPanels>
        </TabGroup>
      </div>
    </div>
  </main>
</template>
