export const formatCurrency = (amount) => {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
    }).format(amount);
};

export const getFullStateName = (fullState, append = false) => {
    const state = ["", "ALMOST FULL", "FULL"][fullState];

    if (append && state) return " (" + state + ")";
    return state;
};
