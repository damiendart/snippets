// A fix for iOS/iPadOS WebKit's non-standard bubbling of mouse events.
//
// For more information about the following fix, see
// <https://www.quirksmode.org/blog/archives/2014/02/mouse_event_bub.html>.
//
// Copyright (C) 2022 Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

// eslint-disable-next-line no-underscore-dangle
function _noop() {
  return () => {};
}

function applyFix() {
  if ('ontouchstart' in document.documentElement) {
    // "document.body.children" returns a "HTMLCollection", which does
    // not have a "forEach" method despite being an array-like object.
    Array.from(document.body.children).forEach(
      (element) => {
        element.addEventListener('mouseover', _noop);
      },
    );
  }
}

function removeFix() {
  Array.from(document.body.children).forEach(
    (element) => {
      element.removeEventListener('mouseover', _noop);
    },
  );
}

export { applyFix, removeFix };
