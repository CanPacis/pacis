document.addEventListener("alpine:init", () => {
  Alpine.directive("check", (el) => {
    const disabled = el.getAttribute("aria-disabled") === "true";
    const input = el.querySelector("input");
    if (!input || disabled) return;

    el.addEventListener("keyup", (e) => {
      if (e.key === " " || e.key === "Enter") {
        input.checked = !input.checked;
      }
    });
  });

  Alpine.store("tabs", {
    init() {
      Alpine.effect(() => {
        const instances = this.value.map((tab) => tab.owner);
        instances.forEach((owner) => {
          const url = new URL(window.location.href);
          const tab = url.searchParams.get(owner);
          if (tab) {
            const index = this.value.findIndex((tab) => tab.owner === owner);
            this.value[index].active = tab;
          }
        });
      });
    },
    value: [],
    define(owner, defaultValue) {
      this.value.push({ owner: owner, active: defaultValue ?? null });
    },
    set(owner, id) {
      const tab = this.value.find((tab) => tab.owner === owner);
      if (!tab) return;
      tab.active = id;
      const url = new URL(window.location.href);
      url.searchParams.set(owner, id);
      history.pushState({}, "", url);

      const tabContent = document.getElementById(`tab-content-${owner}-${id}`);
      if (tabContent) {
        tabContent.focus();
      }
    },
    isActive(owner, id) {
      const tab = this.value.find((tab) => tab.owner === owner);
      return tab?.active === id;
    },
  });
});
