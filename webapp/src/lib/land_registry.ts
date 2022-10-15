export type ID = string;

export type LandRegistry = {
    address: string;
    owners: string[];
    area: number;
    mortgage?: Mortgage;
};

export type Owner = {
    firstName: string;
    lastName: string;
    pesel: string;
    share: number;
};

export type Mortgage = {
    creditor: string;
    amount: number;
};