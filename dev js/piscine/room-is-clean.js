function isMyRoomClean(check) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (check) {
          resolve("My room is clean");
        } else {
          reject("My room is not clean");
        }
      }, 1000); // 1 second delay
    });
  }