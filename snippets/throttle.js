// A simple throttling function.
//
// Copyright (C) 2022 Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

function throttle(callback, interval = 0) {
  let throttled = true;

  return (...args) => {
    if (!throttled) {
      return;
    }

    throttled = false;
    callback.apply(this, args);

    if (interval > 0) {
      // eslint-disable-next-line no-return-assign
      setTimeout(() => throttled = true, interval);
    } else {
      // eslint-disable-next-line no-return-assign
      window.requestAnimationFrame(() => throttled = true);
    }
  };
}

export default throttle;
