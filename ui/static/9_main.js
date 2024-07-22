document.addEventListener("alpine:init", () => {
  Alpine.directive("fire", (el, { expression }, { evaluate }) => {
    const disabled = el.getAttribute("aria-disabled") === "true";
    if (disabled) return;

    el.addEventListener("keyup", (e) => {
      if (e.key === " " || e.key === "Enter") {
        e.preventDefault();
        if (expression) {
          evaluate(expression);
        } else {
          const input = el.querySelector("input");
          input.checked = !input.checked;
        }
      }
    });
  });

  Alpine.directive(
    "click-outside",
    (el, { expression }, { evaluateLater, evaluate, effect, cleanup }) => {
      const condition = el.getAttribute("data-click-outside-if");

      const handler = (e) => {
        if (!e.composedPath().includes(el)) {
          evaluate(expression);
        }
      };

      if (condition) {
        const getCondition = evaluateLater(condition);

        effect(() => {
          getCondition(async (value) => {
            if (value) {
              await Alpine.nextTick();
              document.addEventListener("click", handler);
            } else {
              document.removeEventListener("click", handler);
            }
          });
        });
      } else {
        document.addEventListener("click", handler);
        cleanup(() => {
          document.removeEventListener("click", handler);
        });
      }
    }
  );

  Alpine.directive("focusable", (el, { expression }) => {
    const child = el.children[0];
    if (!child) return;

    child.classList.add("focus-visible:focus-ring");
    child.classList.add("focusable");
    child.setAttribute("tabindex", "0");
    if (expression) {
      child.setAttribute("x-fire", expression);
    }
  });

  Alpine.directive("child-on", (el, { value, modifiers, expression }) => {
    const child = el.children[0];
    if (!child) return;

    // todo: handle modifiers
    child.setAttribute(`x-on:${value}`, expression);
  });

  Alpine.directive("menu", (el, _, { evaluateLater, evaluate, effect }) => {
    const getIsOpen = evaluateLater(el.getAttribute("data-is-menu-open"));
    const closeMenu = el.getAttribute("data-close-menu");
    const autofocus = el.getAttribute("data-autofocus");

    const closeHandler = (e) => {
      if (!e.composedPath().includes(el)) {
        evaluate(closeMenu);
      }
    };

    const trapFocus = () => {
      const items = tabbable
        .focusable(el)
        .filter((el) => el.getAttribute("aria-hidden"));
      if (items.length === 0) {
        return;
      }

      let currentFocusIndex = 0;
      const getNext = () => {
        currentFocusIndex = (currentFocusIndex + 1) % items.length;
        return items[currentFocusIndex];
      };
      const getPrev = () => {
        currentFocusIndex =
          currentFocusIndex - 1 < 0 ? items.length - 1 : currentFocusIndex - 1;
        return items[currentFocusIndex];
      };

      if (autofocus) {
        items[currentFocusIndex].focus();
      }

      const trapHandler = (e) => {
        if (e.key === "Tab") {
          e.preventDefault();

          if (e.shiftKey) {
            getPrev().focus();
          } else {
            getNext().focus();
          }
        } else if (e.key === "Escape") {
          evaluate(closeMenu);
        }
      };

      document.addEventListener("keydown", trapHandler);

      for (const item of items) {
        item.addEventListener("keydown", (e) => {
          switch (e.key) {
            case "ArrowDown":
              e.preventDefault();
              getNext().focus();
              break;
            case "ArrowUp":
              e.preventDefault();
              getPrev().focus();
          }
        });
      }

      return () => {
        document.removeEventListener("keydown", trapHandler);
      };
    };

    let cleanupTrapFocus = null;

    effect(() => {
      getIsOpen(async (isOpen) => {
        if (isOpen) {
          await Alpine.nextTick();
          document.addEventListener("click", closeHandler);
          cleanupTrapFocus = trapFocus();
        } else {
          document.removeEventListener("click", closeHandler);
          if (cleanupTrapFocus) {
            cleanupTrapFocus();
            cleanupTrapFocus = null;
          }
        }
      });
    });
  });

  const fuzzy = new window.uFuzzy({});

  Alpine.directive(
    "autocomplete",
    (el, { expression }, { evaluate, evaluateLater, effect }) => {
      const id = evaluate(expression);
      const input = el.querySelector("input[type='search']");
      const data = JSON.parse(document.getElementById(id).textContent);

      input.addEventListener("input", (e) => {});

      const valueGetter = evaluateLater("value");

      effect(() => {
        valueGetter((value) => {
          if (value.trim().length === 0) {
            evaluate("open = false");
            return;
          }

          const found = fuzzy
            .filter(
              data.map((d) => d.label),
              value
            )
            .map((d) => data[d]);

          const options = el.querySelectorAll(`[data-id]`);

          if (found.length > 0) {
            for (const option of options) {
              if (found.some((f) => f.id === option.getAttribute("data-id"))) {
                option.setAttribute("x-show", "true");
                option.setAttribute("aria-hidden", "true");
              } else {
                option.setAttribute("x-show", "false");
                option.removeAttribute("aria-hidden");
              }
            }

            evaluate("open = true");
          }
        });
      });
    }
  );

  Alpine.directive("snippet", (el, { expression }, { evaluate }) => {
    const value = evaluate(expression);

    el.addEventListener("click", () => {
      navigator.clipboard.writeText(value);
      evaluate("done = true");
      setTimeout(() => {
        evaluate("done = false");
      }, 1000);
    });
  });

  Alpine.magic("copy", () => {
    return (data) => {
      navigator.clipboard.writeText(data);
    };
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
