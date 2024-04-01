import { getDownloadURL } from 'firebase/storage'

import { initializeApp } from 'firebase/app'
import { getStorage, ref, uploadBytesResumable } from 'firebase/storage'

const firebaseConfig = {
  apiKey: 'AIzaSyBVemItgmLJkcSg1J-JUfyOakBYbQDuVR0',
  authDomain: 'service-30af5.firebaseapp.com',
  projectId: 'service-30af5',
  storageBucket: 'service-30af5.appspot.com',
  messagingSenderId: '40139711073',
  appId: '1:40139711073:web:5780cfbde109c00f080020',
}

const app = initializeApp(firebaseConfig)
const storage = getStorage(app)

/**
 * Uploads a file to Firebase Storage
 * @param {File} file The file to upload
 * @return {Promise<string>} A promise that resolves to the download URL
 */
export async function uploadImages(file) {
  const storageRef = ref(storage, 'images/' + file.name)

  const uploadTask = uploadBytesResumable(storageRef, file, {
    contentType: file.type,
  })

  //   Listen for state changes, errors, and completion of the upload.
  //   uploadTask.on(
  //     'state_changed',
  //     (snapshot) => {
  //       // Get task progress, including the number of bytes uploaded and the total number of bytes to be uploaded
  //       const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100
  //       console.log('Upload is ' + progress + '% done')

  //       switch (snapshot.state) {
  //         case 'paused':
  //           console.log('Upload is paused')
  //           break
  //         case 'running':
  //           console.log('Upload is running')
  //           break
  //       }
  //     },
  //     (error) => {
  //       // A full list of error codes is available at
  //       // https://firebase.google.com/docs/storage/web/handle-errors
  //       switch (error.code) {
  //         case 'storage/unauthorized':
  //           // User doesn't have permission to access the object
  //           break

  //         case 'storage/canceled':
  //           // User canceled the upload
  //           break

  //         // ...

  //         case 'storage/unknown':
  //           // Unknown error occurred, inspect error.serverResponse
  //           break
  //       }
  //     },
  //     () => {
  //       // Upload completed successfully, now we can get the download URL
  //       getDownloadURL(uploadTask.snapshot.ref).then((url) => {
  //         downloadURL = url
  //         console.log('File available at', downloadURL)
  //         return downloadURL
  //       })
  //     }
  //   )

  const url = uploadTask
    .then((snapshot) => {
      const url = getDownloadURL(snapshot.ref).then((url) => {
        console.log('File available at', url)
        return url
      })
      return url
    })
    .catch((error) => {
      console.log(error)
    })
  return url
}
