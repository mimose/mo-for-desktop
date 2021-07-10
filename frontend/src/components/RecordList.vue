<template>
  <v-container fluid>
    <div style="position: absolute; width: 100%; z-index: 2">
      <v-autocomplete
        v-model="selectedSpace"
        class="pl-14 pr-6"
        :items="spaces"
        item-text="name"
        item-value="key"
        :loading="spaceLoading? `primary` : `false`"
        :disabled="spaceLoading"
        chips
        solo
        dense
        deletable-chips
        label="选择你的空间"
        @change="changeSelectedSpace"
      ></v-autocomplete>
      <!-- <v-select
        v-model="selectedSpace"
        class="pl-14 pr-6"
        :menu-props="{bottom: true, offsetY: true}"
        :items="spaces"
        item-text="name"
        item-value="key"
        :loading="spaceLoading? `primary` : `false`"
        :disabled="spaceLoading"
        chips
        solo
        dense
        deletable-chips
        label="选择你的空间"
        no-data-text="default"
        @change="changeSelectedSpace"
      ></v-select> -->
    </div>

    <div
      v-if="recordGroup != null && recordGroup.length > 0"
      class="py-13 pl-14 pr-6"
      style="position: absolute;height: 90%; width:100%; overflow: auto; z-index: 1"
    >

      <v-tabs v-model="tabModel" centered height=50 fixed-tabs slider-color="accent darken-1">
        <v-tab
          v-for="group in recordGroup"
          :key="group.groupKey"
          :href="`#tab-${group.groupKey}`"
        >
          {{ group.groupName }}


          <v-progress-circular
            :value="progress(group.groupKey)"
            class="ml-4"
            size="20"
            width="10"
          ></v-progress-circular>
        </v-tab>
      </v-tabs>

      <v-tabs-items style="background: none" v-model="tabModel">
        <v-tab-item
          v-for="group in recordGroup"
          :key="group.groupKey"
          :value="`tab-${group.groupKey}`"
        >
          <!-- <v-card flat> -->
            <!-- <v-card-actions> -->
              <v-expansion-panels
                v-if="group.groupValues.length > 0"
                accordion
                multiple
                focusable
              >
                <template v-for="(record, i) in group.groupValues">
                  <v-expansion-panel
                    :key="`${i}-${record.key}`"
                    :readonly="record.readonly"
                  >
                    <v-expansion-panel-header disable-icon-rotate style="padding: 5px 10px 0px 10px">
                      <v-switch
                        dense
                        inset
                        v-model="record.done"
                        @change="check(record)"
                        style="max-width: 10%"
                      >
                      </v-switch>
                      <span
                        :class="
                          (record.done && 'grey--text') || 'primary--text'
                        "
                        class="ml-4"
                        v-text="record.title"
                      ></span>
                      <template v-slot:actions>
                        <v-icon dense v-if="record.done" color="success">
                          mdi-check
                        </v-icon>
                        <v-icon dense v-if="!record.done" color="primary"></v-icon>
                      </template>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content eager>
                      {{record.content}}
                    </v-expansion-panel-content>
                  </v-expansion-panel>
                </template>
              </v-expansion-panels>
            <!-- </v-card-actions> -->
          <!-- </v-card> -->
        </v-tab-item>
      </v-tabs-items>

      <!--  -->
    </div>

    <div
      class="pl-14 pr-6"
      style="bottom: 5px; position: absolute; width: 100%; z-index: 2"
    >
      <v-bottom-sheet v-model="sheet" persistent eager>
        <template v-slot:activator="{ on, attrs }">
          <!-- <v-btn width="100%" color="accent darken-1" v-bind="attrs" v-on="on">
            新建任务
          </v-btn> -->
          <div class="text-center">
            <v-btn
              width="100%"
              class="my-3"
              color="accent darken-2"
              v-bind="attrs" v-on="on"
              text
            >
              <v-icon dark>
                mdi-plus
              </v-icon>
              新建任务
            </v-btn>
          </div>
        </template>
        <v-sheet class="text-center" height="200px">
          <v-btn class="mt-6" text color="error" @click="sheet = !sheet">
            close
          </v-btn>
          <div class="py-3">
            
          </div>
        </v-sheet>
      </v-bottom-sheet>
    </div>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    selectedSpace: "",
    spaceLoading: false,
    tabModel: "tab-notice",
    spaces: [],
    recordGroup: [
      {
        groupKey: "notice",
        groupName: "notice",
        groupValues: [],
      },
      {
        groupKey: "note",
        groupName: "note",
        groupValues: [],
      },
    ],
    newNoticeTask: null,
    sheet: false,
  }),
  computed: {
    completedRecordCount() {
      return function(groupKey) {
        if(groupKey == 'notice') {
          return this.recordGroup[0].groupValues.filter((task) => task.done).length;
        } else {
          return this.recordGroup[1].groupValues.filter((task) => task.done).length;
        }
      }
    },
    recordCount() {
      return function(groupKey) {
        if(groupKey == 'notice') {
          return this.recordGroup[0].groupValues.length;
        } else {
          return this.recordGroup[1].groupValues.length;
        }
      }
    },
    remainingRecordCount() {
      return function(groupKey) {
        return this.recordCount(groupKey) - this.completedRecordCount(groupKey);
      }
    },
    progress() {
      return function(groupKey) {
        return (this.completedRecordCount(groupKey) / this.recordCount(groupKey)) * 100;
      }
    },
  },
  mounted: function() {
    // 获取空间列表
    this.loadSpaces();
  },
  methods: {
    /**
     * 获取空间列表
     */
    loadSpaces: function() {
      let vm = this;
      vm.spaceLoading = true;
      window.backend.Mo.ListSpaces().then(resp => {
        if(resp.code === 200 && resp.data != null) {
          vm.spaces = resp.data;
          vm.selectedSpace = vm.spaces[0].key
          vm.changeSelectedSpace();
        }
        vm.spaceLoading = false;
      });
    },
    async changeSelectedSpace() {
      let vm = this;
      // 进行全局遮罩，在查询时进行
      vm.$emit('update:appOverlaySync', true)
      await vm.completeRecordGroup('notice', 0);
      await vm.completeRecordGroup('note', 1);
      vm.$emit('update:appOverlaySync', false)
    },
    async completeRecordGroup(groupKey, recordKey) {
      let vm = this;
      let noticeRecordGroup = vm.recordGroup.find(item => item.groupKey === groupKey);
      if(noticeRecordGroup == null) {
        return;
      }
      noticeRecordGroup.groupValues = [];
      if(vm.selectedSpace === '' || vm.selectedSpace === null) {
        return;
      }
      let noticesResp = await window.backend.Mo.ListRecord(vm.selectedSpace, recordKey);
      if(noticesResp.code === 200) {
        if(noticesResp.data != null) {
          noticesResp.data.forEach(element => {
            console.log(element)
            let value = {
              key: element.key,
              done: element.done,
              title: element.title,
              content: element.content,
              readonly: element.content == null || element.content == ''
            }
            noticeRecordGroup.groupValues.push(value);
          });
        }
      }
    },
    create() {
      this.recordGroup[0].groupValues.push({
        done: false,
        text: this.newNoticeTask,
      });

      this.newNoticeTask = null;
    },
    check(record) {
        console.log(record)
    }
  },
};
</script>