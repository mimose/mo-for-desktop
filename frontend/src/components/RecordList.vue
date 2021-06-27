<template>
    <v-container fluid>
        <v-select
            v-model="selectedGroup"
            class="pl-14"
            :items="items"
            item-text="name"
            item-value="key"
            chips
            solo
            deletable-chips
            label="Choose one group"
            no-data-text="default"
            @change="clearSelectedGroup"
        ></v-select>
    </v-container>
</template>

<script>
export default {
    data: () => ({
        selectedGroup: '',
        items:[{
            key: 1,
            name: "first"
        },{
            key: 2,
            name: "second"
        },{
            key: 3,
            name: "three"
        },],
    }),
    methods: {
        clearSelectedGroup() {
            
        }
    },
};
</script>