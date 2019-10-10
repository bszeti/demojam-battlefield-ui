rm -rf ./static

#yarn build from react dir
cd react
yarn 
yarn build

#move app to backend directory
mv build ../static

cd ..



