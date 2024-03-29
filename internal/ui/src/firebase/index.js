import { initializeApp } from 'firebase/app'
import { getStorage, ref, uploadBytes } from 'firebase/storage'

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
const storageRef = ref(storage, 'images/')

/**
 * Uploads an image to storage.
 *
 * @param {...File} files - The image files to upload
 * @return {Promise} Promise that resolves on upload complete
 */
export function uploadImage(...files) {
  const downloadLink = uploadBytes
    .apply(storageRef, files)
    .then((snapshot) => {
      console.log('Uploaded a blob or file!')
      // For the image link
      const publicUrl = snapshot.ref.getDownloadURL().toString()
      console.log(publicUrl)
      return publicUrl
    })
    .catch((err) => console.error(err))

  // For displaying in the view
  return downloadLink
}
