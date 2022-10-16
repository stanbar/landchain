export type Individual = {
    name: string;
    pesel: string;
    publicKey: string;
};

export type Owner = Individual;

export type Mortgage = {
    creditor: string;
    amount: number;
};

export type NotarialAct = {
    landID: string;
    notary: Individual;
    creationDate: string;
    description: string;
    owner: Individual;
};

export type Land = {
    id: string;
    address: string;
    owner: string;
    area: number;
    mortgage: Mortgage;
    notarialAct: NotarialAct;
};

export enum State {
    SUBMITTED,
    BLOCKED,
    REJECTED,
    ACCEPTED,
    WAITING_FOR_PAYMENT,
    WAITING_FOR_APPROVAL
};

export type Request = {
    landID: string;
    creationTime: string;
    parties: Individual[];
    intermediary: Individual;
    newOwner: Owner;
    previousNotarialAct: NotarialAct;
    newNotarialAct: NotarialAct;
};

export type SignedRequest = {
    id: string;
    signatures: string[];
    state: State;
} & Request;