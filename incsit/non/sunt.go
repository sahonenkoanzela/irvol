signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
if err != nil {
    log.Fatal(err)
}
