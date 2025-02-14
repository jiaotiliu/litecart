<template>
  <MainLayouts>
    <div v-if="!route.params.page_slug">
      <header>
        <h1>Pages</h1>
        <div>
          <FormButton type="submit" name="New" color="green" ico="arrow-right" @click="openDrawer(null, 'add')" />
        </div>
      </header>

      <div class="mx-auto pb-16" v-if="pages.length > 0">
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th class="w-32">Position</th>
              <th class="w-32">Slug</th>
              <th class="w-32">Created</th>
              <th class="w-32">Updated</th>
              <th class="w-24 px-4 py-2"></th>
            </tr>
          </thead>
          <tbody>
            <tr :class="{ 'opacity-30': !item.active }" v-for="(item, index) in pages" :key="item.id">
              <td @click="openPage(item.slug)">{{ item.name }}</td>
              <td @click="openPage(item.slug)">{{ item.position }}</td>
              <td><a :href="`/${item.slug}`" target="_blank">{{ item.slug }}</a></td>
              <td @click="openPage(item.slug)">
                {{ formatDate(item.created) }}
              </td>
              <td v-if="item.updated">{{ formatDate(item.updated) }}</td>
              <td v-else></td>
              <td class="px-4 py-2">
                <div class="flex">
                  <div class="pr-3">
                    <SvgIcon name="pencil-square" class="h-5 w-5" @click="openDrawer(index, 'update')" stroke="currentColor" />
                  </div>
                  <div class="pr-3">
                    <SvgIcon name="rocket" class="h-5 w-5" @click="openDrawer(index, 'seo')" stroke="currentColor" />
                  </div>
                  <div>
                    <SvgIcon :name="item.active ? 'eye' : 'eye-slash'" class="h-5 w-5" @click="updatePageActive(index)" stroke="currentColor" />
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="mx-auto" v-else>Not found pages</div>
    </div>
    <div v-else>
      <header>
        <h1>{{ page.name }}</h1>
      </header>
      <div>
        <Editor v-model="content" />

        <hr class="my-5" />
        <FormButton type="submit" name="Save" color="green" @click="updatePageContent" />
      </div>
    </div>

    <drawer :is-open="isDrawer.open" max-width="710px" @close="closeDrawer">
      <PageAdd    :page="page" :pages="pages" :close="closeDrawer" v-if="isDrawer.action === 'add'" />
      <PageUpdate :page="page" :pages="pages" :close="closeDrawer" v-if="isDrawer.action === 'update'" />
      <ProjectSeo :page="page" :close="closeDrawer" v-if="isDrawer.action === 'seo'" />
    </drawer>
  </MainLayouts>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import MainLayouts from "@/layouts/Main.vue";
import Editor from "@/components/Editor.vue";
import FormButton from "@/components/form/Button.vue";
import Drawer from "@/components/Drawer.vue";
import PageAdd from "@/components/page/Add.vue";
import PageUpdate from "@/components/page/Update.vue";
import ProjectSeo from "@/components/page/Seo.vue";
import { formatDate } from "@/utils/";
import { showMessage } from "@/utils/message";
import { apiGet, apiUpdate } from "@/utils/api";

const route = useRoute();
const router = useRouter();
const pages = ref([]);

const page = ref({});
const content = ref();

const isDrawer = ref({
  open: false,
  action: null,
});

onMounted(() => {
  pagesList();

  if (route.params.page_slug) {
    pageContent(route.params.page_slug);
  }
});

const pagesList = async () => {
  apiGet(`/api/_/pages`).then(res => {
    if (res.success) {
      pages.value = res.result;
    }
  });
};

const openPage = (slug) => {
  router.push({ name: "pagesArticle", params: { page_slug: slug } });
  pageContent(slug);
};

const pageContent = async (slug) => {
  apiGet(`/api/pages/${slug}`).then(res => {
    if (res.success) {
      page.value = res.result;
      content.value = res.result.content;
    } else {
      router.push({ name: "404" });
    }
  });
};

const updatePageContent = async () => {
  apiUpdate(`/api/_/pages/${page.value.id}/content`, page.value).then(res => {
    if (res.success) {
      showMessage(res.message);
    } else {
      showMessage(res.result, "connextError");
    }
  })
};

const updatePageActive = async (index) => {
  apiUpdate(`/api/_/pages/${pages.value[index].id}/active`, null).then(res => {
    if (res.success) {
      pages.value[index].active = !pages.value[index].active;
    }
  });
};

const openDrawer = (index, action) => {
  isDrawer.value.open = true;
  isDrawer.value.action = action;
  page.value = pages.value[index]
};

const closeDrawer = () => {
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};
</script>
