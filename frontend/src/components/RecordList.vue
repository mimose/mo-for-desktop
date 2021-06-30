<template>
    <v-container fluid>
        <v-select
            v-model="selectedSpace"
            class="pl-14"
            :items="items"
            item-text="name"
            item-value="key"
            chips
            solo
            deletable-chips
            label="Choose one space"
            no-data-text="default"
            @change="clearSelectedSpace"
        ></v-select>
    </v-container>
</template>

<script>
export default {
    data: () => ({
        selectedSpace: '',
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
        clearSelectedSpace() {
            
        }
    },
};
</script>