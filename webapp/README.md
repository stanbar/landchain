# create-svelte

Everything you need to build a Svelte project, powered by [`create-svelte`](https://github.com/sveltejs/kit/tree/master/packages/create-svelte).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```bash
# create a new project in the current directory
npm create svelte@latest

# create a new project in my-app
npm create svelte@latest my-app
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.

## Views

1. Create Request
1.1. Create request object.
1.2. Add button "Add signature", which adds signature to the request.
1.3. On backend, call `chaincode.SubmitRequest(signedRequest)`

2. View All Waiting For Approval Requests
On backend, call `chaincode.GetAllRequestsWaitingForApproval()`
2.1 Add two buttons
- Accept. Call on backend `chaincode.AcceptRequest(request.ID)`
- Reject. Call on backend `chaincode.RejectRequest(request.ID)`

3. View All Lands
On backend, call `chaincode.GetAllLands()`