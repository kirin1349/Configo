import { Num2 } from "./Num2";

export class Num3 extends Num2
{
    z: number = 0;

    constructor(x: number = 0, y: number = 0, z: number = 0)
    {
        super(x, y);

        this.z = z;
    }
}